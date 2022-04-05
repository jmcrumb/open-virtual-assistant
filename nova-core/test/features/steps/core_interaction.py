import io
import subprocess
import sys
from time import sleep
from unittest.mock import patch

from behave import *
from core.nova_core import NovaCore
from core.plugin_registry import registry
from main import (audio_response_handler, main, text_only_input_handler,
                  text_only_response_handler)
from nlp.nlp import SpeechRecognition

def after_all(context):
    context.stdout_mock.truncate(0)
    context.stdout_mock.seek(0)

@given('nova is running')
def step_impl(context):
    context.stdout_mock = patch('sys.stdout', new_callable=io.StringIO).start()
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
    print(output)
    if text not in output:
        fail('%r not in %r' % (text, output))

@when('wait "{text}" seconds')
def step_impl(context, text):
    sleep(float(text))

@then(u'the output is successfully converted to audio')
def step_impl(context):
    raise NotImplementedError(u'STEP: the output is successfully converted to audio')

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
    # context.control_loop = subprocess.run(['python', 'main.py'])
    pass


@when(u'we request the plugin "{plugin}" download via CLI')
def step_impl(context, plugin):
    raise NotImplementedError(u'STEP: When we request the plugin "placeholder" download via CLI')


@then(u'the plugin "{plugin}" is in the plugin registry')
def step_impl(context, plugin):
    registered_plugins = [plugin.__name__ for plugin in registry.keys()]
    if not plugin in registered_plugins:
        fail(f'{plugin} not in registry')
