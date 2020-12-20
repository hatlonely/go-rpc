#!/usr/bin/env bash

function List() {
    find tpl -type d -depth 1 | awk '{print(substr($0, 5, length($0)))}'
}

function Render() {
    mkdir -p tmp
    tpl=$1
    cfg=$2
    gomplate -f "tpl/${tpl}/environment.sh.tpl" -c .="${cfg}" > "tmp/environment.sh" && echo "render success"
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
