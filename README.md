[![Build Status](https://cloud.drone.io/api/badges/raohwork/marionette-go/status.svg)](https://cloud.drone.io/raohwork/marionette-go)
[![GoDoc](https://godoc.org/github.com/raohwork/marionette-go?status.svg)](https://godoc.org/github.com/raohwork/marionette-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/raohwork/marionette-go)](https://goreportcard.com/report/github.com/raohwork/marionette-go)

# What is marionette 

1. A marionette is a puppet controlled from above using wires or strings depending on regional variations. (Wikipedia)
2. Marionette is the remote protocol that lets out-of-process programs communicate with, instrument, and control Gecko-based browsers. (Mozilla Source Tree Docs)

# Synopsis

```go
package main

import (
    marionette "github.com/raohwork/marionette-go"
    "github.com/raohwork/marionette-go/mnclient"
    "github.com/raohwork/marionette-go/mnsender"
)

func main() {
    s, err := mnsender.NewTCPSender("127.0.0.1:2828", 0)
    // handler error here

    if err := s.Start(); err != nil {
        // handle error
    }
    defer s.Stop()

    cl := &mnclient.Commander{Sender: s}

    // go to google
    cl.Navigate("https://www.google.com")
}
```

# Helper Docker image

There's a Docker image hosted at [Docker Hub](https://hub.docker.com/r/ronmi/go-firefox).
It is a helper for me (and projects using marionette-go, maybe) to run tests at local machine or docker based CI env like [DroneCI](https://drone.io).

The image accepts 3 envvars

- `GO_VER`: Go version to download, like `1.11.5`.
- `FX_VER`: Firefox version to download, like `65.0.1`.
- `XVFB`: Disable headless mode, use xvfb instead.

Refer `dev-test.sh` and `.drone.yaml` for example usage of this image.

# Pitfalls

- I am still learning how actions works. It does not work as expected now. I will
  fix it sometime, maybe next major version as data structures will doubtlessly
  need redesigning.
- The behavier of screenshot commands varies with versions and headless mode.
  Since it still "take screenshot on the document/viewport/element", I will not 
  write workarounds about it.
- The essential command "NewWindow" is not implemented in Firefox 65 (and lower).
- I am not native English speaker so some sentences in docs may look weird. Feel 
  free to submit issue/PR about that.

# Promise of compatibility

From the first stable version (3.0.0), new versions will strictly following these
rules.

### bugfix version (last digit)

- Fixes bug.

Codes depends on older/newer version will work unchanged (if not depends on the buggy behavier I fixed).

### minor version (mid digit)

- New feature is introduced.
- Marks some features to be deprecated.
- Alone with some bug fixes.

Codes depends on older version will work unchanged (if not depends on the buggy behavier I fixed).

### major version (first digit)

- Remove deprecated feature in previous major version.
- Mozilla changes the command semantically in new version Firefox.
- Mozilla introduces new command in new version Firefox.
- Mozilla removes the support to some command in new version Firefox.
- Marionette protocol version changed.
- Refactor/rewrite the code.

Your code will work differently or fail to compile.

There will be a release note to address the changes and lowest supported Firefox version.

# License

LGPLv3 | MPL2.0
