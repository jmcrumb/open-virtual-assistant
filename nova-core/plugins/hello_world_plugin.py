from core.abstract_plugin import NovaPlugin

class HelloWorldPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return ['hello', 'hi', 'howdy', 'hello there']

    def execute(self, command: str) -> str:
        return 'Hello! My name is Nova'

    def help_command(self, command: str) -> str:
        return 'Help the world'