import sys
import unittest
from urllib import response
from core.nova_core import NovaCore
from nlp.nlp import SpeechRecognition, UNKNOWN_VALUE_ERROR

def main():

    # Sets text only CLI interface
    TEXT_DEBUG = False

    if TEXT_DEBUG:
        input_handler = text_only_input_handler
        response_handler = text_only_response_handler
    else:
        # Set voice [OPTIONAL]
        sr: SpeechRecognition = SpeechRecognition()
        sr.set_voice(17)
        input_handler = microphone_input_handler
        response_handler = audio_response_handler
    core = NovaCore(response_handler)

    # Input loop
    while True:
        try:
            command: str = input_handler()
            core.invoke(command)
        except UNKNOWN_VALUE_ERROR:
            print('[Speech to text] Unknown Value')

def text_only_response_handler(response: str):
    print(f'\nNova> {response}\nUser> ', end='')

def text_only_input_handler() -> str:
    return input('User> ')

def microphone_input_handler() -> str:
    sr: SpeechRecognition = SpeechRecognition()
    return sr.speech_to_text()

def audio_response_handler(response: str):
    sr: SpeechRecognition = SpeechRecognition()
    sr.text_to_speech(response)

if __name__ =='__main__':
    main()
