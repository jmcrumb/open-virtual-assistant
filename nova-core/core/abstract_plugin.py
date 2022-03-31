from abc import ABC, abstractclassmethod

class NovaPlugin(ABC):

    def __init__(self):
        pass
    
    @abstractclassmethod
    def get_keywords(self) -> list:
        pass

    @abstractclassmethod
    def execute(self, command: str)-> str:
        pass

    @abstractclassmethod
    def execute_secondary_command(self, command: str) -> str:
        pass

    def request_history(self) -> list:
        return []
    
    def __string__(self) -> str: 
        return __class__.__name__