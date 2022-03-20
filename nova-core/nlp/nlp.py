from threading import Semaphore

import speech_recognition as sr
import pyttsx3

class SpeechRecognition:
    '''Singleton implmentation of Speech Recognition utility functions'''

    _instance = None

    def __new__(cls):
        if cls._instance is None:
            cls._instance = super(SpeechRecognition, cls).__new__(cls)
            cls._instance._initialization_routine()
        return cls._instance

    def _initialization_routine(self) -> None:
        self.recognizer = sr.Recognizer()
        self.tts_engine = pyttsx3.init()
        self._tts_mutex = Semaphore()


    def set_voice(self, voice_index: int) -> None:
        voices = self.tts_engine.getProperty('voices')
        self.tts_engine.setProperty('voice', voices[voice_index].id)

    def speech_to_text(self, input_: str) -> str:
        '''Speech to text which accepts str or .wav'''
        return input_
  
    
    def speech_to_text(self) -> str:
         with sr.Microphone() as source:
            self.recognizer.adjust_for_ambient_noise(source, duration=0.2)
            print('Say something')
            audio_data = self.recognizer.listen(source)
            print('processing')
            text: str = self.recognizer.recognize_google(audio_data)
            print(text)
            return text
    

    def text_to_speech(self, input_: str, language='en') -> any:
        self._tts_mutex.acquire()
        # if self.tts_engine._inLoop:
        #     self.tts_engine.endLoop()
        self.tts_engine.say(input_)
        self.tts_engine.runAndWait()
        self.tts_engine.endLoop()   # add this line
        self.tts_engine.stop()
        self._tts_mutex.release()