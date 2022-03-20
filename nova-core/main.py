import signal
from core.nova_core import NovaCore
from nlp.nlp import SpeechRecognition

def main():
    # Sets text only CLI interface
    TEXT_DEBUG = False

    if TEXT_DEBUG:
        core = NovaCore(text_only_response_handler)
    else:
        core = NovaCore(audio_response_handler)

    # Set voice [OPTIONAL]
    sr: SpeechRecognition = SpeechRecognition()
    sr.set_voice(17)

    # Input loop
    while True:
        if TEXT_DEBUG:
            command: str = text_only_input_handler()
        else:
            command: str = microphone_input_handler()
        core.invoke(command)

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
