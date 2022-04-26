from abc import ABC, abstractclassmethod

"""
NovaPlugin is an abstract class that models the construction of any plugin in for the Nova System.

Plugins are singleton with the vanilla Nova Core. When a plugin is called, it will hold priority
until the user gives a command that is not handled within the plugin. This command will cause 
Nova Core to search for another plugin with matching keywords or provide a "command not understood"
error.
"""
class NovaPlugin(ABC):

    def __init__(self):
        """
        __init__() is the constructor of the plugin.

        :param self: represents the instance of the class
        :return: an instance of the class

        Use this function to set class variables for any saved information the plugin needs.
        Primarily, these values are the state of the plugin and settings, but can be used anyway
        you'd use class variables on another project.
        """
        pass
    
    @abstractclassmethod
    def get_keywords(self) -> list:
        """
        get_keywords() returns a list of designated keywords that correspond to this plugin alone.

        :param self: represents the instance of the class
        :return: a list of keywords that relate to this plugin

        Try to be unique but intuitive, though conflicts are easily solved by editing this function 
        in the plugins causing the issue.
        """
        pass

    @abstractclassmethod
    def execute(self, command: str)-> str:
        """
        execute() is what is called by the core to handle the command given by the user.

        :param self: represents the instance of the class
        :param command: a string containing the plain text command from the user
        :return: a string response related to the command

        This method can range from a simple logic statement to a state machine so be open with 
        experiementing. It is recommended that you use helper functions to support this main 
        method.
        """
        pass

    @abstractclassmethod
    def help_command(self, command: str) -> str:
        """
        help_command() is a required helper function that should be called by execute_command()
        if the user provides an unrecognizable command.

        :param self: represents the instance of the class
        :param command: a string containing the plain text command from the user
        :return: a string response that provides help corresponding to the command

        It can be as simple as a preset string with information to a state machine that allows for 
        multiple responses.
        """
        pass

    def request_history(self) -> list:
        """
        request_history() is TODO (ignore)

        :param self: represents the instance of the class
        :return:
        """
        return []
    
    def __str__(self) -> str:
        """
        __str__() is represents the class objects as a string

        :param self: represents the instance of the class
        :return: a string representing the class
        """
        return __class__.__name__