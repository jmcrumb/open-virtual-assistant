from core.abstract_plugin import NovaPlugin

class JokePlugin(NovaPlugin):
	def __init__(self):
		self.jokes = [
			"What do you call a person with no arms or legs? ... a cripple!",
			"Three men walk into a bar, the first one gets a drink, the second sees someone he knows, and the third immediately sexually harasses a girl and is arrested the end."
		]
		self.last = 0

	def get_keywords(self) -> list:
		return ['joke', 'funny']

	def execute(self, command: str) -> str:
		self.last += 1
		if self.last > len(self.jokes):
			self.last = 0

		return self.jokes[self.last]

	def help_command(self, command: str) -> str:
		return ('This joke plugin is powered by the pure comedic genius of Kanan Boubion.'
            ' It is an application that spreads joy and laughter to all who use it.'
            ' To hear a joke, just say:'
            ' Tell me a joke. or. Say something funny.')