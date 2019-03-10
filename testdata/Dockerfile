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

FROM debian:stable-slim

# install firefox-esr for runtime dependencies only, not gonna use it.
RUN apt-get update \
 && apt-get upgrade -y \
 && apt-get install -y --no-install-recommends \
        wget ca-certificates firefox-esr \
        build-essential bzip2 git xvfb \
 && apt-get clean -y \
 && rm -fr /var/lib/apt/lists/*

ADD *.sh /usr/local/bin/

ENV GOPATH=/go
ENTRYPOINT ["/usr/local/bin/run-test.sh"]
