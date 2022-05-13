from core.abstract_plugin import NovaPlugin
import wikipediaapi

class WikiPlugin(NovaPlugin):

    def get_keywords(self) -> list:
        return ['wiki']

    def execute(self, command: str) -> str:
        if ('wiki' in command):
            wiki_wiki = wikipediaapi.Wikipedia('en')

            item = command[command.rfind("wiki")+4:].strip()
            page = wiki_wiki.page(item)
            if page.exists():
                msg = (page.summary[0:page.summary.find('.', page.summary.find('.')+1)])
            else:
                msg = ('{} not found. Please try again.'.format(item))
        else:
            msg = None
        
        return msg

    def help_command(self, command: str) -> str:
        return 'This plugin returns the wiki information to the user.'