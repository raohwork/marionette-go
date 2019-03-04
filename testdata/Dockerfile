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
