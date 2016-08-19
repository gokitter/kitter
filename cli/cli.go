package main

import (
	"fmt"
	"os"

	"github.com/gokitter/kitter/client"
	"github.com/gokitter/kitter/server"
)

func main() {
	// Set up a connection to the server.
	if os.Args[2] == "server" {
		server.StartRPCServer(os.Args[1])
	} else if os.Args[2] == "receive" {
		c := clientfactory.Create(os.Args[1])
		c.ReadStream(&callback{})
	} else {
		c := clientfactory.Create(os.Args[1])
		c.WriteMessage(os.Args[3])
	}
}

type callback struct{}

func (c *callback) NewMessage(message string) {
	fmt.Println(message)
}
