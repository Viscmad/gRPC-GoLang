package main

import (
	"fmt"
	"log"
	"net"

	pb "gRPC-GoLang/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type todoServer struct {
	pb.TodoServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTodoServiceServer(grpcServer, &todoServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server %v", err)
	}
	fmt.Println("Starting Server!")
}
