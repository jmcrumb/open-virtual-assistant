import time

from core.nova_core import NovaCore
from nlp.nlp import UNKNOWN_VALUE_ERROR, SpeechRecognition


def get_env() -> dict:
    return {
        # Name of virtual assistant
        'NAME': 'nova',

         # Dictates how many seconds can pass since the last question before Nova considers it 
        # a new conversation
        'ATTENTION_SPAN': 15.0,
        'last_invoked': 0.0,

        # How long Nova listens for their name before releasing resources to response thread
        'INVOCATION_TIMEOUT': 2.0
    }

def main():

    sr: SpeechRecognition = SpeechRecognition()

    # Sets text only CLI interface
    TEXT_DEBUG: bool = False

    env: dict = get_env()


    if TEXT_DEBUG:
        input_handler = text_only_input_handler
        response_handler = text_only_response_handler
    else:
        # Set voice [OPTIONAL]
        # sr.set_voice(17)
        input_handler = microphone_input_handler
        response_handler = audio_response_handler
    core = NovaCore(response_handler)

    # Input loop
    while True:
        try:
            command: str = input_handler(env)
            core.invoke(input_=command)
        except UNKNOWN_VALUE_ERROR:
            core.invoke(input_=None, unknown_input=True)

def is_ongoing_conversation(last_invoked: float, attention_span: float) -> bool:
    return (time.time() - last_invoked) < attention_span

def text_only_response_handler(response: str):
    print(f'\nNova> {response}\nUser> ', end='')

def text_only_input_handler(env: dict) -> str:
    return input('User> ')

def microphone_input_handler(env: dict) -> str:
    sr: SpeechRecognition = SpeechRecognition()
    if is_ongoing_conversation(env['ATTENTION_SPAN'], env['last_invoked']) or sr.is_nova_invocation(
            keyword=env['NAME'],
            timeout=env['INVOCATION_TIMEOUT']
        ):
        print('[NOVA] Invoked')
        output: str = sr.speech_to_text(timeout=env['ATTENTION_SPAN'])
        env['last_invoked'] = time.time()
        return output
    return None

def audio_response_handler(response: str):
    sr: SpeechRecognition = SpeechRecognition()
    sr.text_to_speech(response)

if __name__ =='__main__':
    main()
