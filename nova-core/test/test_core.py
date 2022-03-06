from io import StringIO
from time import sleep
import unittest
from unittest.mock import patch
from main import response_handler
from core.nova_core import NovaCore

# CLI: python -m unittest discover

class NovaCoreTests(unittest.TestCase):

    def setUp(self):
        self.core = NovaCore(response_handler)

    def test_plugin_initialization(self):
        self.assertIn('hello', self.core.syntax_tree.root.keys())

    def test_invoke(self):
        with patch('sys.stdout', new = StringIO()) as fake_out:
            self.core.invoke('hello')
            sleep(1)
            self.assertTrue('Hello! My name is Nova.' in fake_out.getvalue())

    def test_invoke_secondary_command(self):
        with patch('sys.stdout', new = StringIO()) as fake_out:
            self.core.invoke('hello')
            sleep(1)
            self.core.invoke('nice to meet you')
            sleep(1)
            self.assertTrue('To teach me more fun things to do, go to the plugin store.' in fake_out.getvalue())

if __name__ == '__main__':
    unittest.main()