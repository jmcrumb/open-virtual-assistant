from core.abstract_plugin import NovaPlugin

class HelloWorldPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return ['hello', 'hi', 'hey']

    def execute(self, command: str) -> str:
        if ('hello' in command, 'hi' in command or 'hey' in command):
            return 'Hello! My name is Nova. How can I help you?'
        
        return None

    def help_command(self, command: str) -> str:
        return 'Helping the world by helping you.'