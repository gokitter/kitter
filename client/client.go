package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/gokitter/kitter/kitter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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

	client := kitter.NewKitterClient(conn)

	if os.Args[1] == "receive" {
		readStream(client)
	} else {
		writeMessage(client, os.Args[1])
	}

}

func writeMessage(client kitter.KitterClient, message string) {
	client.Keet(context.Background(), &kitter.Message{From: "Nic", Content: message})
}

func readStream(client kitter.KitterClient) {
	stream, err := client.KeetStream(context.Background(), &kitter.Filter{})
	if err != nil {
		grpclog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		grpclog.Println(message.Content)
	}
}
