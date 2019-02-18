// Package marionette defines basic/shared data to connect to Firefox marionette server
//
// Example usage
//
//     package main
//
//     import (
//         marionette "github.com/raohwork/marionette-go"
//         "github.com/raohwork/marionette-go/ito"
//         "github.com/raohwork/marionette-go/shirogane"
//     )
//
//     func main() {
//         client := &shirogane.Mixed{} // use default address
//         go client.Start()
//
//         client.Sync(&ito.NewSession{})
//         client.Sync(&ito.Navigate{URL: "https://google.com"})
//
//         // get title using marionette service
//         cmdTitle := &ito.GetTitle{}
//         msg, _ := client.Sync(cmdTitle)
//         title, _ := cmdTitle.Decode(msg)
//         log.Print("Page title: ", title)
//
//         // get title using javascript
//         cmdJS := &ito.ExecuteScript{
//             Script: `return document.title`,
//             Args: []interface{}{},
//         }
//         msg, _ = client.Sync(cmdJS)
//         var jsTitle string
//         cmdJS.Decode(msg, &jsTitle)
//         log.Print("Page title: ", jsTitle)
//     }
package marionette // import "github.com/raohwork/marionette-go"
