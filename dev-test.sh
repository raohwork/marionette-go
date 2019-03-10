#!/bin/bash

# This file is part of marionette-go
#
# marionette-go is distributed in two licenses: The Mozilla Public License,
# v. 2.0 and the GNU Lesser Public License.
#
# marionette-go is distributed in the hope that it will be useful, but WITHOUT
# ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
# FOR A PARTICULAR PURPOSE.
#
# See License.txt for further information.


DIR="/go/src/github.com/raohwork/marionette-go/v3"
mkdir -p .cache/opt

function runtest {
    echo ''
    echo ''
    echo "Test ${3} against Go ${1} Firefox ${2}"
    GOV="$1"
    FXV="$2"
    PKG="$3"
    shift; shift; shift
    docker run --rm \
           -v "$(pwd):${DIR}" \
           -v "$(pwd)/.cache/opt:/opt" \
           -v "$(pwd)/.cache:/home/user" \
           -e "HOME=/home/user" \
           -e "GO_VER=${GOV}" \
           -e "FX_VER=${FXV}" \
           --workdir "$DIR" \
           --user "${UID}:${GID}" \
           ronmi/go-firefox \
           go test -p 2 -bench . -benchmem "$PKG" "$@"
}

CNT=0
MAX=3
if [[ $DEBUG != "" ]]
then
    MAX=1
fi

function jobc {
    CNT=$((CNT+1))

    if [[ $CNT -ge $MAX ]]
    then
        wait
        CNT=0
    fi
}

set -e

for go in 1.10.8 1.11.5 1.12
do
    for fx in 66.0b9 66.0b12
    do
        for pkg in . ./mnsender ./mnclient ./tabmgr
        do
            runtest $go $fx $pkg &
            jobc
        done
    done
done

for go in 1.10.8 1.11.5 1.12
do
    for fx in 64.0 65.0
    do
        echo "Test tabmgr against Go ${go} Firefox ${fx}"
        runtest $go $fx ./tabmgr -run TestTabManager &
        jobc
    done
done

wait
