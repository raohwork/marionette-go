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

gopkg="https://dl.google.com/go/go${GO_VER}.linux-amd64.tar.gz"
fxpkg="https://download.mozilla.org/?product=firefox-${FX_VER}&lang=en-US&os=linux64"
godir="/opt/go/${GO_VER}"
fxdir="/opt/fx/${FX_VER}"

# install go
mkdir -p "$godir"
if [[ ! -x "${godir}/go/bin/go" ]]
then
    wget -q -O - "$gopkg" | tar zxvf - -C "$godir"
fi

# install fx
mkdir -p "$fxdir"
if [[ ! -x "${fxdir}/firefox" ]]
then
    wget -q -O - "$fxpkg" | tar jxvf - -C "$fxdir" --strip-components=1
fi
