package main

import (
	"log"
	"os"

	"github.com/gokitter/kitter/client"
	"github.com/gokitter/kitter/kitter"
	"github.com/gokitter/kitter/server"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	if os.Args[1] == "server" {
		server.StartRPCServer()
	} else if os.Args[1] == "receive" {
		c := kitter.NewKitterClient(conn)
		client.ReadStream(c)
	} else {
		c := kitter.NewKitterClient(conn)
		client.WriteMessage(c, os.Args[1])
	}

}
