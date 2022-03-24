import io
from threading import Semaphore

import gtts
import speech_recognition as sr
from playsound import playsound

UNKNOWN_VALUE_ERROR = sr.UnknownValueError

UNKNOWN_VALUE_ERROR = sr.UnknownValueError

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
        # self.tts_engine = pyttsx3.init()
        self._mutex = Semaphore()
        # pygame.mixer.init()

    # def set_voice(self, voice_index: int) -> None:
    #     voices = self.tts_engine.getProperty('voices')
    #     self.tts_engine.setProperty('voice', voices[voice_index].id)

    def speech_to_text(self, input_: str) -> str:
        '''Speech to text which accepts str'''
        return input_
  
    
    def speech_to_text(self) -> str:
         with sr.Microphone() as source:
            self.recognizer.adjust_for_ambient_noise(source, duration=0.5)
            print('Say something')
            audio_data = self.recognizer.listen(source)
            print('processing')
            text: str = self.recognizer.recognize_google(audio_data)
            print(text)
            return text
    
    def speech_to_text(self, timeout=None) -> str:
        text: str = ''
        with sr.Microphone() as source:
            self._mutex.acquire()
            try:
                self.recognizer.adjust_for_ambient_noise(source, duration=3)
                print('[NLP] Listening')
                audio_data = self.recognizer.listen(source, timeout=timeout)
                print('[NLP] Processing')
                text = self.recognizer.recognize_google(audio_data)
                print(f'[NLP] Input recognized: {text}')
            finally:
                self._mutex.release()
                return text


    def text_to_speech(self, input_: str, language='en', tld='com') -> any:
        with io.BytesIO() as f:
            self._mutex.acquire()
            # self.tts_engine.say(input_)
            # self.tts_engine.runAndWait()
            gtts.gTTS(input_, lang=language).save('.temp_tts_output.mp3')
            playsound('.temp_tts_output.mp3')
            # gtts.gTTS(input_, lang=language, tld=tld).write_to_fp(f)
            # f.seek(0)
            # song = AudioSegment.from_file(f, format="mp3")
            # play(song)
            self._mutex.release()

    def is_nova_invocation(self, keyword='nova', timeout=2) -> bool:
        text = self.speech_to_text(timeout=timeout).lower()
        if keyword in text: 
            print(f'[NLP] Invocation recognized')
            return True
        return False
