package main

import (
	"context"
	"log"
	"net"
	"versioning-go-grpc-service/greetings_v1"
	"versioning-go-grpc-service/greetings_v2"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type serverV1 struct {
	greetings_v1.UnimplementedGreetingsServer
}

type serverV2 struct {
	greetings_v2.UnimplementedGreetingsServer
}

// SayHello implements helloworld.GreeterServer
func (s *serverV1) SayHello(ctx context.Context, in *greetings_v1.HelloRequest) (*greetings_v1.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &greetings_v1.HelloResponse{Message: "Hello " + in.GetName()}, nil
}

// SayHello implements helloworldv2.GreeterServer
func (s *serverV2) SayHello(ctx context.Context, in *greetings_v2.HelloRequest) (*greetings_v2.HelloResponse, error) {
	log.Printf("V2 Received: %v", in.GetName(), in.GetLastName())
	return &greetings_v2.HelloResponse{Message: "Hello " + in.GetName() + in.GetLastName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetings_v1.RegisterGreetingsServer(s, &serverV1{})
	greetings_v2.RegisterGreetingsServer(s, &serverV2{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
