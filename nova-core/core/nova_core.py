import json
from queue import Queue
from threading import Semaphore, Thread

from nlp import nlp
from plugins.command_not_found_plugin import CommandNotFoundPlugin
from plugins.hello_world_plugin import HelloWorldPlugin

from core.abstract_plugin import NovaPlugin

import core.plugin_registry as plugin_registry

NOT_FOUND_PLUGIN = CommandNotFoundPlugin()

class SyntaxTree:

    def __new__(cls):
        if not hasattr(cls, 'instance'):
            cls.instance = super(SyntaxTree, cls).__new__(cls)
            cls.instance._initialize()
        return cls.instance

    def _initialize(self):
        self.root: dict = {}
        self.not_found: NovaPlugin = NOT_FOUND_PLUGIN
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
        self.hits: dict = {}
        max_depth: int = 0

        #for token in tokenized_command:
        for i in range(len(tokenized_command)):
            token = tokenized_command[i]
            if token in branch:
                depth: int = self.match_command_rec(branch, tokenized_command[i:], 0)
                if depth > max_depth: max_depth = depth

        return self.hits[max_depth][0] if len(self.hits) > 0 else self.not_found

    def match_command_rec(self, branch: dict, tokenized_command: list, depth: int) -> int:
        if len(tokenized_command) > 0 and tokenized_command[0] in branch:
            return self.match_command_rec(branch[tokenized_command[0]], tokenized_command[1:], depth + 1)
        else:
            if not (depth in self.hits): self.hits[depth] = []
            if 'plugin' in branch:
                self.hits[depth].append(branch['plugin'])
            return depth

class AsyncPluginThreadManager:

    def __init__(self, response_handler, CommandNotFound):
        self.active_threads: set = set()
        self.command_not_found = CommandNotFound
        self.mru_plugin: NovaPlugin = None
        self.syntax_tree = SyntaxTree()

        self.buffer: Queue = Queue()
        self.io_mutex: Semaphore = Semaphore()
        self.keep_alive = True
        self.response_thread = ResponseLoop(self, response_handler)

        self.response_thread.start()

    def dispatch(self, command: str, plugin=None, try_again=True):
        if not plugin:
            if not self.mru_plugin:
                plugin = self.syntax_tree.match_command(command)
            else:
                plugin = self.mru_plugin
        t: Thread = PluginThread(self, plugin, command, try_again=try_again)
        t.start()
        self.active_threads.add(t)

    def __del__(self):
        self.keep_alive = False
        self.response_thread.join()  

class PluginThread(Thread):

        def __init__(self, manager: AsyncPluginThreadManager, plugin: NovaPlugin, command: str, try_again=True):
            super().__init__()
            self.manager: AsyncPluginThreadManager = manager
            self.plugin: NovaCore = plugin
            self.command: str = command
            self.try_again: bool = try_again

        def run(self):
            response = self.plugin.execute(self.command)
            if response:
                if self.plugin != self.manager.command_not_found:
                    self.manager.mru_plugin = self.plugin
                self.send_response(response)
            elif self.try_again:
                best_fit_plugin: NovaPlugin = self.manager.syntax_tree.match_command(self.command)
                if best_fit_plugin != self.manager.command_not_found:
                    self.manager.mru_plugin = best_fit_plugin
                self.manager.dispatch(self.command, best_fit_plugin, try_again=False)
            else:
                self.manager.dispatch(self.command, self.manager.command_not_found)

        def send_response(self, response):
            self.manager.buffer.put(response)
            self.manager.active_threads.discard(self)

class ResponseLoop(Thread):

    def __init__(self, manager: AsyncPluginThreadManager, response_handler):
        super().__init__()
        self.manager = manager
        self.response_handler = response_handler

    def run(self):
        manager = self.manager
        while manager.keep_alive:
            self.response_handler(manager.buffer.get())
            manager.buffer.task_done()

class NovaCore:

    def __init__(self, response_handler):
        self.plugins: list[NovaPlugin] = []

        for Plugin in plugin_registry.registry:
            self.plugins.append(Plugin())

        self.syntax_tree: SyntaxTree = SyntaxTree()
        self._initialize_plugins()
        self.thread_manager = AsyncPluginThreadManager(response_handler, NOT_FOUND_PLUGIN)

    def _initialize_plugins(self):
        for plugin in self.plugins:
            self.syntax_tree.add_plugin(plugin)

    def invoke(self, input_=None, unknown_input=False):
        if unknown_input:
            self.thread_manager.dispatch(self.thread_manager.command_not_found, '')
            return
        if not input_: 
            return
            
        command: str = input_.lower()

        self.thread_manager.dispatch(command)