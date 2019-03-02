// This file is part of marionette-go
//
// marionette-go is free software: you can redistribute it and/or modify it
// under the terms of the GNU Lesser General Public License as published by the
// Free Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// marionette-go is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.
//
// You should have received a copy of the GNU Lesser General Public License along
// with marionette-go. If not, see <https://www.gnu.org/licenses/>.

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
//         s, err := mnsender.NewSender("127.0.0.1:2828", 0)
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
package marionette // import "github.com/raohwork/marionette-go"
