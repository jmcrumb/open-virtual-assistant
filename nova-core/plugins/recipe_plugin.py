from json import loads
from core.abstract_plugin import NovaPlugin
from requests import Response, get


class RecipePlugin(NovaPlugin):

    def __init__(self):
        self.api_key = '39c940f961484513854eff8866bc6482'
        
    
    def get_keywords(self) -> list:
        return ['recipe for', 'get recipe for', 'cuisine for']

    def execute(self, command: str) -> str:
        if 'recipe' in command:
            return self._get_recipe(command)
        elif 'cuisine' in command:
            return self._get_recipe(command)

    def _call_api(self, base: str, args: str) -> Response:
        return get(
            f'{base}?{args}',
            headers={'x-api-key': self.api_key}
            )

    def _get_recipe(self, command: str) -> str:
        query = command.replace(' ', '+')
        rsp = self._call_api(
            'https://api.spoonacular.com/food/converse',
            f'text={query}'
        )

        jsRsp = rsp.json()
        if rsp.ok:
            msg = jsRsp['answerText']
            for recipe in jsRsp['media']:
                msg += f'{recipe["title"]}.  '
        else:
            msg = ('There seems to be an issue with the weather API. A {} error as recieved.\n'
                'Error Code: {}'.format(rsp.status_code, jsRsp['error']['code']))

        return msg

    def _get_cuisine(self, command: str) -> str:
        return 'work in progress'

    def help_command(self, command: str) -> str:
        """
        help_command() is a required helper function that should be called by execute_command()
        if the user provides an unrecognizable command.

        :param self: represents the instance of the class
        :param command: a string containing the plain text command from the user
        :return: a string response that provides help corresponding to the command

        It can be as simple as a preset string with information to a state machine that allows for 
        multiple responses.
        """
        pass
