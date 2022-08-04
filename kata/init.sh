#!/usr/bin/env bash
# vim: autoindent tabstop=4 shiftwidth=4 expandtab softtabstop=4 ft=sh

set -e # Exit on any error. Use `COMMAND || true` to nullify
set -E # Functions inherit error trap
set -u # Error on unset variables. Use ${var:-alternate value} to bypass
set -f # Error on attempted file globs (e.g. *.txt )
set -o pipefail # Failed commands in pipes cause the whole pipe to fail
shopt -s nocasematch

if [[ ${#} -eq 0 ]]; then
    >&2 echo "You must pass in a directory name"
    exit 1
fi

dir=$( echo "${@}" | tr ' A-Z' '-a-z' )

if [[ $dir =~ ^[1-9]- ]]; then
    dir="0$dir"
fi

mkdir "$dir"
cd "$dir"

go mod init github.com/oliverisaac/${PWD##*/git/}

git add .
git commit . -n -m "🎉 Initial setup of: $dir"

