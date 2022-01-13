package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc/greeting"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9080, "The server port")
)

type server struct {
	greeting.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *greeting.HelloRequest) (*greeting.HelloReply, error) {
	log.Printf("Received : %s", in.GetName())
	return &greeting.HelloReply{
		Message: "hello " + in.GetName(),
	}, nil
}

func main() {
	flag.Parse()
	lis, er := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if er != nil {
		log.Fatalf("failed to listen: %v", er)
	}

	s := grpc.NewServer()
	greeting.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v",
		lis.Addr())
	if er := s.Serve(lis); er != nil {
		log.Fatalf("failed to serve: %v", er)
	}
}
