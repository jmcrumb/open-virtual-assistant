from io import StringIO
from time import sleep
import unittest
from unittest.mock import patch
from main import text_only_response_handler
from core.nova_core import NovaCore, SyntaxTree, AsyncPluginThreadManager
from plugins.command_not_found_plugin import CommandNotFoundPlugin
from plugins.hello_world_plugin import HelloWorldPlugin
from core.abstract_plugin import NovaPlugin


class NovaCoreTests(unittest.TestCase):

    def setUp(self):
        self.core = NovaCore(text_only_response_handler)

    def test_plugin_initialization(self):
        self.assertIn('hello', self.core.syntax_tree.root.keys())

    def test_invoke(self):
        with patch('sys.stdout', new = StringIO()) as fake_out:
            self.core.invoke('hello')
            sleep(1)
            self.assertTrue('Hello! My name is Nova' in fake_out.getvalue())

    def tearDown(self):
        self.core.syntax_tree.root = {}

class SyntaxTreeTests(unittest.TestCase):
    
    def setUp(self):
        self.syntax_tree: SyntaxTree = SyntaxTree()
        self.syntax_tree.add_plugin(HelloWorldPlugin())

    def test_add_plugin(self):
        self.assertIn('hello', self.syntax_tree.root.keys())

    def test_match_command(self):
        plugin: NovaPlugin = self.syntax_tree.match_command('hello')
        self.assertTrue(isinstance(plugin, HelloWorldPlugin))

    def test_match_command_multilevel(self):
        plugin: NovaPlugin = self.syntax_tree.match_command('hello there')
        self.assertTrue(isinstance(plugin, HelloWorldPlugin))

    def tearDown(self):
        self.syntax_tree.root = {}  
        
class AsyncPluginThreadManagerTests(unittest.TestCase):
    
    def setUp(self):
        self.thread_manager: AsyncPluginThreadManager = AsyncPluginThreadManager(
            text_only_response_handler,
            CommandNotFoundPlugin
        )

    def test_dispatch_and_recieve(self):
        hw_plugin: HelloWorldPlugin = HelloWorldPlugin()
        with patch('sys.stdout', new = StringIO()) as fake_out:
            self.thread_manager.dispatch('hello', hw_plugin)
            sleep(1)
            self.assertTrue('Hello! My name is Nova' in fake_out.getvalue(), msg=f'output was {fake_out.getvalue()}')   

if __name__ == '__main__':
    unittest.main()