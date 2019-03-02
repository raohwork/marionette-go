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
    s, err := mnsender.NewSender("127.0.0.1:2828", 0)
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

LGPLv3
