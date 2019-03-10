// This file is part of marionette-go
//
// marionette-go is distributed in two licenses: The Mozilla Public License,
// v. 2.0 and the GNU Lesser Public License.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE.
//
// See License.txt for further information.

// Package marionette defines basic/shared data to connect to Firefox marionette server
//
// Example usage
//
//     package main
//
//     import (
//         marionette "github.com/raohwork/marionette-go"
//         "github.com/raohwork/marionette-go/mnclient"
//         "github.com/raohwork/marionette-go/mnsender"
//     )
//
//     func main() {
//         s, err := mnsender.NewTCPSender("127.0.0.1:2828", 0)
//         // handler error here
//
//         if err := s.Start(); err != nil {
//             // handle error
//         }
//         defer s.Stop()
//
//         cl := &mnclient.Commander{Sender: s}
//
//         // go to google
//         cl.Navigate("https://www.google.com")
//     }
package marionette // import "github.com/raohwork/marionette-go/v3"
