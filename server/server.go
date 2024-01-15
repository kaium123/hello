package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gg/hello"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServiceServer
}

func (s *server) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Res: "Hello " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		fmt.Println(err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &server{})

	log.Println("Server is listening on port 50051")
	if err := s.Serve(listener); err != nil {
		fmt.Println(err)
	}
}
