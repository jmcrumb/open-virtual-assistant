from time import sleep
from behave import *
from core.nova_core import NovaCore
from main import text_only_response_handler, text_only_input_handler
import sys
import io

def before_all(context):
    context.real_stdout = sys.stdout
    context.stdout_mock = io.StringIO()
    sys.stdout = context.stdout_mock

def after_all(context):
    sys.stdout = context.real_stdout

@given('nova is running')
def step_impl(context):
    context.core = NovaCore(text_only_response_handler)
    
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