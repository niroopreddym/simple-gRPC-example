package main

import (
	"context"
	"fmt"

	"github.com/niroopreddym/simple-gRPC-example/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8790", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)
	a := 10
	b := 20
	req := &proto.Request{A: int64(a), B: int64(b)}
	response, err := client.Multiply(context.Background(), req)
	fmt.Println(response, err)
}
