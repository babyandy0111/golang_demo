package main

import (
	"context"
	pb "golang_demo/proto/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) SaySomeThing(ctx context.Context, in *pb.AndyTestRequest) (*pb.AndyTestReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.AndyTestReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAndyServerServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
