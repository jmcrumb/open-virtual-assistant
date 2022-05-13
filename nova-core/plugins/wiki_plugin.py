from core.abstract_plugin import NovaPlugin
import wikipediaapi

class WikiPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return ['wiki']

    def execute(self, command: str) -> str:
        if ('wiki' in command):
            wiki_wiki = wikipediaapi.Wikipedia('en')

            page_py = wiki_wiki.page('Python_(programming_language)')
            print(page_py.exists())
            return 'It is currently {}.'.format(datetime.now().strftime("%H:%M"))
        
        return None

    def help_command(self, command: str) -> str:
        return 'This plugin returns the current time to the user.'