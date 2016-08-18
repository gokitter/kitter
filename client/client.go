package client

import (
	"context"
	"io"
	"log"

	"github.com/gokitter/kitter/kitter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// KitterCallback allows a method to be called on the native code
type KitterCallback interface {
	NewMessage(message string)
}

// KitterClient creates a new client that can be used for sending and receiving messages
// to a Kitter server
type KitterClient struct {
	client kitter.KitterClient
}

// NewKitterClient creates a new Kitter client and connets to the given server
func NewKitterClient(server string) *KitterClient {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := kitter.NewKitterClient(conn)

	return &KitterClient{client: c}
}

// WriteMessage sends a message to the kitter server
func (k *KitterClient) WriteMessage(message string) {
	k.client.Keet(context.Background(), &kitter.Message{From: "Nic", Content: message})
}

// ReadStream reads a stream of messages from the kitter server.
// When a new message is received the NewMessage method is called on the
// provided callback
func (k *KitterClient) ReadStream(callback KitterCallback) {
	stream, err := k.client.KeetStream(context.Background(), &kitter.Filter{})
	if err != nil {
		grpclog.Fatalf("%v.ListFeatures(_) = _, %v", k.client, err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Fatalf("%v.ListFeatures(_) = _, %v", k.client, err)
		}

		callback.NewMessage(message.Content)
	}
}

// Close closes the connection to the server
func (k *KitterClient) Close() {
	k.Close()
}
