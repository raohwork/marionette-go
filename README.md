# What is marionette 

1. A marionette is a puppet controlled from above using wires or strings depending on regional variations. (Wikipedia)
2. Marionette is the remote protocol that lets out-of-process programs communicate with, instrument, and control Gecko-based browsers. (Mozilla Source Tree Docs)

# WIP project

This library still in early development stage, apis might change in any time.

- [X] Knight Sliver-V: Implement most used commands and simple clients as proof of concept.
- [ ] Grisel: Implement all commands and briefly describe its usage.
- [ ] Gun of Feather: Write feature-rich clients.

# Synopsis

```go
package main

import (
    marionette "github.com/raohwork/marionette-go"
    "github.com/raohwork/marionette-go/ito"
    "github.com/raohwork/marionette-go/shirogane"
)

func main() {
    client := &shirogane.Mixed{} // use default address
    go client.Start()
    
    client.Sync(&ito.NewSession{})
    client.Sync(&ito.Navigate{URL: "https://google.com"})

    // get title using marionette service
    cmdTitle := &ito.GetTitle{}
    msg, _ := client.Sync(cmdTitle)
    title, _ := cmdTitle.Decode(msg)
    log.Print("Page title: ", title)
    
    // get title using javascript
    cmdJS := &ito.ExecuteScript{
        Script: `return document.title`,
        Args: []interface{}{},
    }
    msg, _ = client.Sync(cmdJS)
    var jsTitle string
    cmdJS.Decode(msg, &jsTitle)
    log.Print("Page title: ", jsTitle)
}
```

# License

You should explicitly choose one license from MIT, GPL or LGPL before using this library. Default to LGPLv3 if not claimed explicitly.
