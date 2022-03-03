from functools import lru_cache
from core.abstract_plugin import NovaPlugin
from plugins.command_not_found_plugin import CommandNotFoundPlugin
from plugins.hello_world_plugin import HelloWorldPlugin
from nlp import nlp


class rlist(list):
    '''
    LOCAL USE ONLY
    '''
    
    def __init__(self, default):
        self._default = default
    def __setitem__(self, key, value):
        if key >= len(self):
            self += [self._default] * (key - len(self) + 1)
        super(rlist, self).__setitem__(key, value)


class SyntaxTree:

    def __init__(self, commandNotFoundPlugin: NovaPlugin):
        self.root: dict = {}
        self.not_found: NovaPlugin = commandNotFoundPlugin
        self.hits: rlist = []

    def add_plugin(self, plugin: NovaPlugin):
        keywords: list[str] = plugin.get_keywords()
        branch: dict = self.root

        for keyword in keywords:
            tokenized_keyword: list[str] = keyword.split(' ')
            for i in range(len(tokenized_keyword)):
                if tokenized_keyword[i] not in branch: 
                    branch[tokenized_keyword[i]] = {}
                branch = branch[tokenized_keyword[i]]
            if tokenized_keyword[-1] in branch and 'plugin' in branch[tokenized_keyword[-1]]:
                raise Exception(f'Keyword Conflict. Plugin already exists at the keyword {keyword}')
            branch[tokenized_keyword[-1]] = {'plugin': plugin}

    def match_command(self, command: str) -> NovaPlugin:
        tokenized_command: list[str] = command.split(' ')
        branch: dict = self.root
        self.hits = []
        max_depth: int = 0

        for token in tokenized_command:
            if token in branch:
                depth: int = self.match_command_rec(branch, tokenized_command, 0)
                if depth > max_depth: max_depth = depth
        return self.hits[max_depth] if len(self.hits) > 0 else self.not_found

    # @lru_cache(50, False)
    def match_command_rec(self, branch: dict, tokenized_command: list, depth: int) -> int:
        if len(tokenized_command) > 0 and tokenized_command[0] in branch:
            return self.match_command_rec(branch[tokenized_command[0]], tokenized_command[1:], depth + 1)
        else:
            if self.hits[depth] is None: self.hits[depth] = []
            self.hits[depth].append(branch['plugin'])
            return depth


class NovaCore:

    def __init__(self):
        self.plugins: list[NovaPlugin] = [
            HelloWorldPlugin()
        ]
        self.syntax_tree: SyntaxTree = SyntaxTree(CommandNotFoundPlugin())
        self.initialize_plugins()

    def initialize_plugins(self):
        for plugin in self.plugins:
            self.syntax_tree.add_plugin(plugin)

    # TODO: implement multithreading with producer consumer model to handle responses
    def invoke(self, input_):
        command: str = nlp.speech_to_text(input_).lower()
        self.current_plugin = self.query_syntax_tree(command)
        return nlp.text_to_speech(self.current_plugin.execute(command))

    def query_syntax_tree(self, command: str) -> NovaPlugin:
        return self.syntax_tree.match_command(command)



