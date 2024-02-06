package main

import (
	"context"
	hello_grpc "grpc/grpc-mein/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:8888", grpc.WithInsecure())
	defer conn.Close()
	client := hello_grpc.NewHelloServiceClient(conn)
	r, err := client.SayHi(context.Background(), &hello_grpc.Req{Message: "req from client"})
	if err != nil {
		return
	}

	println(r.Message)
}
