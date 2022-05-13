from core.abstract_plugin import NovaPlugin

class JokePlugin(NovaPlugin):
	def __init__(self):
		self.jokes = [
			"What do you call a cow with no legs? ... ground beef!",
			"What do you call a fake noodle? ... an im-pasta!",
			"Why can't you trust atoms? ... they make up everything!",
			"How much does a polar bear weight? ... enough to break the ice!",
			"Why should I tell you a joke? You're the clown here!"
		]
		self.last = 0

	def get_keywords(self) -> list:
		return ['joke', 'funny']

	def execute(self, command: str) -> str:
		if ('joke' in command or 'funny' in command or 'another' in command):
			self.last += 1
			if self.last >= len(self.jokes):
				self.last = 0

			return self.jokes[self.last]
		
		return None

	def help_command(self, command: str) -> str:
		return ('This joke plugin is powered by the pure comedic genius of Kanan Boubion.'
            ' It is an application that spreads joy and laughter to all who use it.'
            ' To hear a joke, just say:'
            ' Tell me a joke. or. Say something funny.')