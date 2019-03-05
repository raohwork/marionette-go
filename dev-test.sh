#!/bin/bash
DIR="/go/src/github.com/raohwork/marionette-go"
mkdir -p .cache/opt

function runtest {
    echo ''
    echo ''
    echo "Test ${3} against Go ${1} Firefox ${2}"
    docker run -it --rm \
           -v "$(pwd):${DIR}" \
           -v "$(pwd)/.cache/opt:/opt" \
           -v "$(pwd)/.cache:/home/user" \
           -e "HOME=/home/user" \
           -e "GO_VER=${1}" \
           -e "FX_VER=${2}" \
           --workdir "$DIR" \
           --user "${UID}:${GID}" \
           ronmi/go-firefox \
           go test -p 2 -bench . -benchmem "$3"
}

set -e

for go in 1.10.8 1.11.5 1.12
do
    for fx in 66.0b9 66.0b12
    do
        for pkg in . ./mnsender ./mnclient ./tabmgr
        do
            runtest $go $fx $pkg
        done
    done
done

for go in 1.10.8 1.11.5 1.12
do
    for fx in 64.0 65.0
    do
        echo "Test tabmgr against Go ${go} Firefox ${fx}"
        runtest $go $fx ./tabmgr
    done
done
