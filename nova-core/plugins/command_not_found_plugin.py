from core.abstract_plugin import NovaPlugin

class CommandNotFoundPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return []

    def execute(self, command: str) -> str:
        return 'I\'m sorry, I don\'t understand.'