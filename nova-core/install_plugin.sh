#!/bin/bash

SRC="/Users/MaxonCrumb1/Documents/DevelopementWorkspaces/OVA/open-virtual-assistant/nova-core"

boot() {
    python "${PWD}/main.py"
}

rebootCore() {
    # TODO: catch kill process in python, kill child threads in deconstructor
    PID=$(ps | grep -i "main.py$" | cut -d " " -f 2)
    if [[ -n "$PID" ]]; then
        kill "$PID"
    fi    

    open -a "Terminal" "install_plugin.sh"
}

installPlugin() {
    pip install $1
    python -m pip freeze >> "${PWD}/requirements.txt"

    echo "$2,$3" >> "plugins.csv"

    # TODO: insulate against injection attack

    PLUGIN_REGISTRY="plugins=[\n\t"

    while IFS="," read -r path className
    do
        PLUGIN_REGISTRY="from $path import $className\n$PLUGIN_REGISTRY\n\t$className,"
    done < <(tail -n +2 plugins.csv)

    echo -e "$PLUGIN_REGISTRY\n]" > "core/plugin_registry.py"

    rebootCore

}

updatePlugins() {
    python -m pip install --upgrade $1
    python -m pip freeze >> "${PWD}/requirements.txt"
    rebootCore
}

cd $SRC
source "${PWD}/venv/bin/activate"

case "$1" in 
    "install") installPlugin $2 $3 $4;;
    "update") updatePlugins $2;;
    "reboot") rebootCore;;
    *) boot;;
esac