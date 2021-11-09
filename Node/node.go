package Node

import (
	"context"
	"fmt"
	"log"
	"net"

	n "tokenring/DISYS_M2"

	"google.golang.org/grpc"
)

type Node struct {
	//include server and client
	client n.TokenRingClient
	server n.TokenRingServer

	ID           int
	Port         string
	NextNodePort string
}

func (node *Node) initNode() {
	lis, err := net.Listen("tcp", node.Port)

	if err != nil {
		log.Fatal("Failed to listen on port %s: %v", node.Port, err)
	}
	grpcServer := grpc.NewServer()
	n.RegisterTokenRingServer(grpcServer, node.server)

	conn, err := grpc.Dial(node.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()

	client := n.NewTokenRingClient(conn)

	context := context.Background()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("ShittyChat server has not started successfully :( %v", err)

	}

	token := n.Token{Access: true, Message: "test", NextToken: 3}
	client.GrantToken(context, &token)

}

func (node *Node) CreateServer() {

	list, err := net.Listen("tcp", ":"+node.Port)
	if err != nil {
		log.Fatal("Failed to listen on port %s: %v", node.Port, err)
	}
	fmt.Println("Listening")

	grpcServer := grpc.NewServer()
	err = grpcServer.Serve(list)
	if err != nil {
		log.Fatal("Failed to start gRPC server: %v", err)
	}
	fmt.Println("Started Server", node.ID)
}

func (node *Node) CreateClient() {
	conn, err := grpc.Dial(node.Port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()

	client := n.NewTokenRingClient(conn)

	context := context.Background()

	out, err := client.GrantToken(context, &n.Token{})
	if err != nil {
		log.Fatal("Failed to open up for publishing chat messages")
	}
}
