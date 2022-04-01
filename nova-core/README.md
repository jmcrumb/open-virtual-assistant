# Nova Core

## Virtual Environment
To set up run `cd nova-core/ && python3 -m venv/ && source venv/bin/activate && pip install - r dependencies.txt` from the root directory.


Enter `source venv/bin/activate` into the CLI upon instantiating a new shell.

## Testing
Enter `python -m unittest discover` to run all unit tests.

## install_plugins.sh script

`./install_plugins.sh boot` Runs Nova Core

`./install_plugins.sh install name python-path class-name` Installs new plugin to Nova Core and reboots service

`./install_plugins.sh update name` Updates a specific plugin and reboots servive

Note: name refers to the published name on pip.  This application relies on pip's import service to pull python packages.

## Running program 
`python main.py`
