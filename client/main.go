package main

import (
	"fmt"
	"log"

	pb "gRPC-GoLang/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

type todoServer struct {
	pb.TodoServiceServer
}

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Client failed to connect with server")
	}
	defer conn.Close()

	fmt.Println("Connected to Server!")

	// client := pb.NewTodoServiceClient(conn)

	// todos := &pb.TodoItems{
	// 	Todos: []string{""}
	// }
}
