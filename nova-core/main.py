from core.nova_core import NovaCore

def main():
    core = NovaCore()
    
    while True:
        command: str = input('User> ')
        response: str = core.invoke(command)
        print(f'Nova> {response}')

if __name__ =='__main__':
    main()