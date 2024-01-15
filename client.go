package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "gg/hello"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := pb.NewServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Hello(ctx, &pb.Request{Name: name})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Res)
}
