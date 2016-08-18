package main

import (
	"log"
	"os"

	"github.com/gokitter/kitter/client"
	"github.com/gokitter/kitter/kitter"
	"github.com/gokitter/kitter/server"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	if os.Args[2] == "server" {
		server.StartRPCServer(os.Args[1])
	} else if os.Args[2] == "receive" {
		c := kitter.NewKitterClient(conn)
		client.ReadStream(c)
	} else {
		c := kitter.NewKitterClient(conn)
		client.WriteMessage(c, os.Args[2])
	}

}
