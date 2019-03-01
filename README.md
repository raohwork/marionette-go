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

# License

You should explicitly choose one license from MIT, GPL or LGPL before using this library. Default to LGPLv3 if not claimed explicitly.
