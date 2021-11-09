package Node

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Node struct {
	ID           int
	Port         string
	NextNodePort string
}

func (node *Node) CreateServer() {
	list, err := net.Listen("tcp", ":"+node.Port)
	if err != nil {
		log.Fatal("Failed to listen on port %s: %v", node.Port, err)
	}

	grpcServer := grpc.NewServer()
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal("Failed to start gRPC server: %v", err)
	}
	fmt.Println("Started Server")
}
