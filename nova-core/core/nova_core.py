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

    def __init__(self, response_handler, CommandNotFound):
        self.active_threads: set = set()
        self.command_not_found = CommandNotFound()

        self.CAPACITY = 10
        self.buffer: list = []
        self.in_index = 0
        self.out_index = 0
        
        self.mutex = Semaphore()
        self.empty = Semaphore(self.CAPACITY)
        self.full = Semaphore(0)

        self.keep_alive = True
        self.response_thread = ResponseLoop(self, response_handler)
        self.response_thread.start()

    def dispatch(self, plugin: NovaPlugin, command: str, is_secondary: bool=False):
        t: Thread = PluginThread(self, plugin, command, is_secondary)
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
            super().__init__()
            self.manager = manager
            self.plugin = plugin
            self.command = command
            self.is_secondary = is_secondary

        def run(self):
            response = None
            if self.is_secondary:
                response = self.plugin.execute_secondary_command(self.command)
                if response:
                    self.send_response(response)
                else:
                    self.send_response(self.manager.command_not_found.execute(self.command))
            else:
                response = self.plugin.execute(self.command)
                if response:
                    self.send_response(response)
                else:
                    self.is_secondary = True
                    self.run()

        def send_response(self, response):
            self.manager.empty.acquire()
            self.manager.mutex.acquire()
            
            self.manager.buffer.append(response)
            self.manager.active_threads.discard(self)
            
            self.manager.mutex.release()
            self.manager.full.release()

class ResponseLoop(Thread):

    def __init__(self, manager: AsyncPluginThreadManager, response_handler):
        super().__init__()
        self.manager = manager
        self.response_handler = response_handler

    def run(self):
        manager = self.manager
        while manager.keep_alive:
            manager.full.acquire()
            manager.mutex.acquire()
            
            self.response_handler(nlp.text_to_speech(manager.buffer.pop(0)))
            
            manager.mutex.release()
            manager.empty.release()


class NovaCore:

    def __init__(self, response_handler):
        self.plugins: list[NovaPlugin] = [
            HelloWorldPlugin()
        ]
        self.CommandNotFound = CommandNotFoundPlugin
        self.syntax_tree: SyntaxTree = SyntaxTree(self.CommandNotFound())
        self._initialize_plugins()
        self.mru_plugin: NovaPlugin = None
        self.thread_manager = AsyncPluginThreadManager(response_handler, self.CommandNotFound)

    def _initialize_plugins(self):
        for plugin in self.plugins:
            self.syntax_tree.add_plugin(plugin)

    def invoke(self, input_):
        command: str = nlp.speech_to_text(input_).lower()

        plugin: NovaPlugin = self.syntax_tree.match_command(command)
        if self.mru_plugin and isinstance(plugin, self.CommandNotFound):
            self.thread_manager.dispatch(self.mru_plugin, command, is_secondary=True)
        else:
            self.thread_manager.dispatch(plugin, command)
            self.mru_plugin = plugin
