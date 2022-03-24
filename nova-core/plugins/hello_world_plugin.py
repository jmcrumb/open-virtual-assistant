from core.abstract_plugin import NovaPlugin

class HelloWorldPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return ['hello', 'hi', 'howdy', 'hello there']

    def execute(self, command: str) -> str:
        return 'Hello! My name is Nova.'

    def execute_secondary_command(self, command: str) -> str:
        return 'To teach me more fun things to do, go to the plugin store.'