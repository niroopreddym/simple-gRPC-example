package main

import (
	"context"
	"net"

	"github.com/niroopreddym/simple-gRPC-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server provides the protobuf server functionality
type Server struct {
}

func main() {
	listener, err := net.Listen("tcp", ":8790")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &Server{})
	reflection.Register(srv)
	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
}

//Add adds the numbers passed in the request
func (server *Server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

//Multiply multiplies the numbers passed in the request
func (server *Server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}
