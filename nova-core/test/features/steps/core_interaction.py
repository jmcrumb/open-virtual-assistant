from time import sleep
from behave import *
from core.nova_core import NovaCore
from main import text_only_response_handler, text_only_input_handler, audio_response_handler, main
import sys
import io
from nlp.nlp import SpeechRecognition

def before_all(context):
    context.real_stdout = sys.stdout
    context.stdout_mock = io.StringIO()
    sys.stdout = context.stdout_mock

def after_all(context):
    sys.stdout = context.real_stdout

@given('nova is running')
def step_impl(context):
    context.core = NovaCore(text_only_response_handler)

@given('nova is running with audio response')
def step_impl(context):
    context.core = NovaCore(audio_response_handler_wrapper)
    
@when('we ask "{text}"')
def step_impl(context, text):
    context.core.invoke(text)

@then(u'output will include "{text}"')
def step_impl(context, text):
    output = context.stdout_mock.getvalue()
    if text not in output:
        fail('%r not in %r' % (text, output))

@when('wait "{text}" seconds')
def step_impl(context, text):
    sleep(float(text))

@then(u'the output is successfully converted to audio')
def step_impl(context):
    pass

def audio_response_handler_wrapper(response) -> bool:
    try:
        audio_response_handler(response)
    except Exception:
        fail('Exception raised in response handler')

@given(u'nlp is running')
def step_impl(context):
    context.sr = SpeechRecognition()


@when(u'we input to the microphone')
def step_impl(context):
    raise NotImplementedError(u'STEP: When we input to the microphone')


@then(u'we recieve a text transcription')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then we recieve a text transcription')


@given(u'main is running')
def step_impl(context):
    context.control_loop = main()


@when(u'we request the plugin "{plugin}" download via CLI')
def step_impl(context):
    raise NotImplementedError(u'STEP: When we request the plugin "placeholder" download via CLI')


@then(u'the plugin "{plugin}" is in the plugin registry')
def step_impl(context):
    raise NotImplementedError(u'STEP: Then the plugin "placeholder" is in the plugin registry')