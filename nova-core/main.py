import signal
from core.nova_core import NovaCore

def response_handler(response: any):
    print(f'\nNova> {response}\nUser> ', end='')

def main():
    core = NovaCore(response_handler)
    # signal.signal(signal.SIG)
    
    while True:
        command: str = input('User> ')
        core.invoke(command)

if __name__ =='__main__':
    main()
