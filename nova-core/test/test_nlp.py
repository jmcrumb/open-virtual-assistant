import unittest
from nlp.nlp import SpeechRecognition


class AsyncPluginThreadManagerTests(unittest.TestCase):
    
    def setUp(self):
        self.sr: SpeechRecognition = SpeechRecognition()

if __name__ == '__main__':
    unittest.main()