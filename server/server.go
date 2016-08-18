package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gokitter/kitter/kitter"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

const port = ":50051"

type broadcaster struct {
	channels []chan kitter.Message
}

func (b *broadcaster) Write(message kitter.Message) {
	for _, channel := range b.channels {
		channel <- message
	}
}

func (b *broadcaster) Listen() chan kitter.Message {
	channel := make(chan kitter.Message)
	b.channels = append(b.channels, channel)

	return channel
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	broadcaster *broadcaster
}

// SayHello implements helloworld.GreeterServer
func (s *server) Keet(ctx context.Context, in *kitter.Message) (*kitter.Error, error) {
	s.broadcaster.Write(*in)

	return &kitter.Error{Code: -1}, nil
}

func (s *server) KeetStream(filter *kitter.Filter, stream kitter.Kitter_KeetStreamServer) error {
	channel := s.broadcaster.Listen()

	for {
		keet := <-channel
		if err := stream.Send(&keet); err != nil {
			return err
		}

	}
}

func StartRPCServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	kitterServer := &server{broadcaster: &broadcaster{}}

	kitter.RegisterKitterServer(rpcServer, kitterServer)

	fmt.Println("Starting server on: ", port)

	rpcServer.Serve(lis)
}

func main() {
	StartRPCServer()
}
