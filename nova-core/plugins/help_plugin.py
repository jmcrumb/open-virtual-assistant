from core.abstract_plugin import NovaPlugin

class HelpPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return ['help', 'confused']

    def execute(self, command: str) -> str:
        return 'I would like to help you, can you further explain the issue?'

    def execute_secondary_command(self, command: str) -> str:
        if 'general' in command:
            rsp = 'Nova is an open source platform designed to empower users'
        elif 'plugin' in command: #possibly regex this
            rsp = 'Nova\'s plugins can be accessed by using their keywords on the virtual assistant. They can also be managed on the Nova app.'
        else: 
            rsp = 'On second thought, you are beyond help.'
        return rsp

    def help_command(self, command: str) -> str:
        return 'The help command is designed to guide you with questions you might have about the Nova Platform'