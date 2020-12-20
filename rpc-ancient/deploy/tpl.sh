#!/usr/bin/env bash

function List() {
    find tpl -maxdepth 1 -name 'environment*' | awk -v FS="." '{if ($1 == "tpl/environment") print $3}'
}

function Render() {
    mkdir -p tmp
    tpl=$1
    cfg=$2
    gomplate -f "tpl/environment.sh.${tpl}.tpl" -c .="${cfg}" > "tmp/environment.sh" && echo "render success"
}

function Help() {
    echo "sh tpl.sh <action> [env]"
    echo "example:"
    echo "  sh tpl.sh ls"
    echo "  sh tpl.sh render prod ~/.gomplate/prod.json"
}

function main() {
    case "$1" in
        "ls") List;;
        "render") Render "$2" "$3";;
        *) Help;;
    esac
}

main "$@"
