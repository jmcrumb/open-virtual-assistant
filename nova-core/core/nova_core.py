from threading import Semaphore, Thread

from nlp import nlp
from plugins.command_not_found_plugin import CommandNotFoundPlugin
from plugins.hello_world_plugin import HelloWorldPlugin

from core.abstract_plugin import NovaPlugin


class SyntaxTree:

    def __init__(self, commandNotFoundPlugin: NovaPlugin):
        self.root: dict = {}
        self.not_found: NovaPlugin = commandNotFoundPlugin
        self.hits: dict = {}

    def add_plugin(self, plugin: NovaPlugin):
        keywords: list[str] = plugin.get_keywords()
        branch: dict = self.root

        for keyword in keywords:
            tokenized_keyword: list[str] = keyword.split(' ')
            for i in range(len(tokenized_keyword) - 1):
                if tokenized_keyword[i] not in branch: 
                    branch[tokenized_keyword[i]] = {}
                branch = branch[tokenized_keyword[i]]
            if tokenized_keyword[-1] in branch and 'plugin' in branch[tokenized_keyword[-1]]:
                raise Exception(f'Keyword Conflict. Plugin already exists at the keyword {keyword}')
            branch[tokenized_keyword[-1]] = {'plugin': plugin}

    def match_command(self, command: str) -> NovaPlugin:
        tokenized_command: list[str] = command.split(' ')
        branch: dict = self.root
        self.hits = {}
        max_depth: int = 0

        for token in tokenized_command:
            if token in branch:
                depth: int = self.match_command_rec(branch, tokenized_command, 0)
                if depth > max_depth: max_depth = depth
        return self.hits[max_depth][0] if len(self.hits) > 0 else self.not_found

    def match_command_rec(self, branch: dict, tokenized_command: list, depth: int) -> int:
        if len(tokenized_command) > 0 and tokenized_command[0] in branch:
            return self.match_command_rec(branch[tokenized_command[0]], tokenized_command[1:], depth + 1)
        else:
            if not (depth in self.hits): self.hits[depth] = []
            self.hits[depth].append(branch['plugin'])
            return depth

class AsyncPluginThreadManager:

    def __init__(self, response_handler):
        self.active_threads: set = set()
        self.response_queue: list = []

        self.CAPACITY = 10
        self.buffer: list = []
        self.in_index = 0
        self.out_index = 0
        
        self.mutex = Semaphore()
        self.empty = Semaphore(self.CAPACITY)
        self.full = Semaphore(0)

        self.keep_alive = True
        self.response_thread = Thread(target=self.recieve, args=(response_handler, self.keep_alive))
        self.response_thread.start()

    def dispatch(self, plugin: NovaPlugin, command: str, is_secondary: bool=False):
        t: Thread = Thread(target=self.thread_payload, args=(plugin, command, is_secondary))
        t.start()
        self.active_threads.add(t)

    def recieve(self, response_handler, keep_alive):
        while keep_alive:
            self.full.acquire()
            self.mutex.acquire()
            
            response_handler(self.buffer.pop(0))
            
            self.mutex.release()
            self.empty.release()
            
    def __del__(self):
        # TODO: kill recieve thread
        self.keep_alive = False
        self.response_thread.join()
            

class PluginThread(Thread):

        def __init__(self, manager: AsyncPluginThreadManager, plugin: NovaPlugin, command: str, is_secondary: bool=False):
            super(self)
            self.manager = manager
            self.plugin = plugin
            self.command = command
            self.is_secondary = is_secondary

        def run(self):
            response = None
            if self.is_secondary:
                response = self.plugin.execute_secondary_command(self.command)
            else:
                response = self.plugin.execute(self.command) 

            self.manager.empty.acquire()
            self.manager.mutex.acquire()
            
            self.manager.buffer.append(response)
            self.manager.active_threads.discard(self)
            
            self.manager.mutex.release()
            self.manager.full.release()

            

class NovaCore:

    def __init__(self, response_handler):
        self.plugins: list[NovaPlugin] = [
            HelloWorldPlugin()
        ]
        self.CommandNotFound = CommandNotFoundPlugin
        self.syntax_tree: SyntaxTree = SyntaxTree(self.CommandNotFound())
        self.initialize_plugins()
        self.current_plugin: NovaPlugin = None
        self.thread_manager = AsyncPluginThreadManager(response_handler)

    def initialize_plugins(self):
        for plugin in self.plugins:
            self.syntax_tree.add_plugin(plugin)

    # TODO: implement multithreading with producer consumer model to handle responses
    def invoke(self, input_):
        command: str = nlp.speech_to_text(input_).lower()

        plugin: NovaPlugin = self.syntax_tree.match_command(command)
        response: str = None
        if isinstance(plugin, self.CommandNotFound):
            response = self.current_plugin.execute_secondary_command(command)
        if response is None:
            self.current_plugin = plugin       
            response = self.current_plugin.execute(command)

        return nlp.text_to_speech(response)

    def execute_aysnc(self, plugin: NovaPlugin, command: str, is_secondary: bool=False):
        return plugin.execute_secondary_command(command) if is_secondary else plugin.execute(command)
