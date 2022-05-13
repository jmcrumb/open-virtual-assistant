from core.abstract_plugin import NovaPlugin
from datetime import datetime

class TimePlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return ['time']

    def execute(self, command: str) -> str:
        if ('what' in command and 'time' in command):
            return 'It is currently {}.'.format(datetime.now().strftime("%H:%M"))
        
        return None

    def help_command(self, command: str) -> str:
        return 'This plugin returns the current time to the user.'