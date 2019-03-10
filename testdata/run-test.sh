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

/usr/local/bin/install-deps.sh

godir="/opt/go/${GO_VER}"
fxdir="/opt/fx/${FX_VER}"

export PATH="${godir}/go/bin:${PATH}"

FX_PROFILE=$(mktemp -d)
FX_HEADLESS="--headless"

if [[ $XVFB != "" ]]
then
    echo "Starting xvfb..."
    Xvfb :99 > /dev/null 2>&1 &
    echo "Wait 3 seconds for Xvfb to start..."
    sleep 3
    export DISPLAY=":99"
    FX_HEADLESS=""
fi



echo "Starting firefox ${FX_VER} ${FX_HEADLESS}..."
"${fxdir}/firefox" --marionette "$FX_HEADLESS" --profile "$FX_PROFILE" > /dev/null 2>&1 &
echo "Wait 5 seconds for firefox to start..."
sleep 5

if [[ $1 != "" ]]
then
    (
        set -x
        eval "$@"
    )
    RET=$?
    
    if [[ $XVFB != "" ]]
    then
        kill %2 > /dev/null 2>&1
    fi
    kill %1 > /dev/null 2>&1
    wait > /dev/null 2>&1
    rm -fr "$FX_PROFILE" > /dev/null 2>&1

    exit $RET
fi
