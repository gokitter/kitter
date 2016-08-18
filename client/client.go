package client

import (
	"context"
	"io"

	"github.com/gokitter/kitter/kitter"
	"google.golang.org/grpc/grpclog"
)

func WriteMessage(client kitter.KitterClient, message string) {
	client.Keet(context.Background(), &kitter.Message{From: "Nic", Content: message})
}

func ReadStream(client kitter.KitterClient) {
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
