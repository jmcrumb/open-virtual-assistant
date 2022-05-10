from core.abstract_plugin import NovaPlugin

class CommandNotFoundPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return []

    def execute(self, command: str) -> str:
        return 'I\'m sorry, I don\'t understand'

    def help_command(self, command: str) -> str:
        return 'This command is used by the core to respond to users when there is an error'