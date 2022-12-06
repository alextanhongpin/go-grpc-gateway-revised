// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/alextanhongpin/go-grpc-gateway/gen/go/protoc/your/service/v1"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

// server is used to implement yourservicev1.YourServiceServer.
type server struct {
	pb.YourServiceServer
}

// Echo implements yourservicev1.YourServiceServer
func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	log.Printf("Received: %v", in.GetValue())
	return &pb.StringMessage{Value: "Hello " + in.GetValue()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYourServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
