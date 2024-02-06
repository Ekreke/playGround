package main

import (
	"context"
	"fmt"
	hello_grpc "grpc/grpc-mein/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello_grpc.UnimplementedHelloServiceServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "Hello " + "response from server"}, nil
}

func main() {
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	hello_grpc.RegisterHelloServiceServer(s, &server{})
	s.Serve(l)
}
