package main

import (
	"context"
	config "grpc_app/common"
	hello "grpc_app/proto"
	"log"
	"net"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

type server struct {
	hello.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hello.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	srv := grpc.NewServer()
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(srv, healthcheck)
	greeter := &server{}
	hello.RegisterGreeterServer(srv, greeter)
	log.Println("Starting RPC server at", config.SERVICE_HELLO_PORT)
	l, err := net.Listen("tcp", config.SERVICE_HELLO_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_HELLO_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}
