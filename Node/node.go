package Node

import (
	"context"
	"log"
	"net"
	"time"
	pb "tokenring/DISYS_M2"

	"google.golang.org/grpc"
)

type Node struct {
	pb.UnimplementedTokenRingServer
	ID           int
	Port         string
	NextNodePort string
}

func (s *Node) GrantToken(ctx context.Context, token *pb.Token) (*pb.Reply, error) {
	log.Printf("Node %d SENT TO NODE %d", token.IdFrom, ConvertPortToId(token.PortTo))
	log.Printf("Current Node ID: %d PORT: %s NEXTPORT: %s", s.ID, s.Port, s.NextNodePort)
	go s.ClientStart(s.NextNodePort)
	return &pb.Reply{Message: "Token Given"}, nil
}

func ServerStart(node Node) {

	for {
		lis, err := net.Listen("tcp", ":"+node.Port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer()
		pb.RegisterTokenRingServer(grpcServer, &node)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %s", err)
		}

	}
}

func (n *Node) AccessCriticalSection() {
	log.Printf("ACCESS CRITICAL SECTION ID: %d", n.ID)
	time.Sleep(time.Second * 1)
}

func (n *Node) ClientStart(nextPort string) {
	n.AccessCriticalSection()

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+nextPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to port %s - %s", n.NextNodePort, err)
	}
	defer conn.Close()
	c := pb.NewTokenRingClient(conn)

	response, err := c.GrantToken(context.Background(), &pb.Token{Message: "Secret Token Access", IdFrom: int32(n.ID), PortTo: n.NextNodePort})
	if err != nil {
		log.Fatalf("Error when calling GrantToken: %s", err)
	}
	log.Printf("Response from server: %s", response.GetMessage())
}

func ConvertPortToId(port string) int {
	if port == "5000" {
		return 0
	}
	if port == "5001" {
		return 1
	}
	if port == "5002" {
		return 2
	}
	return 420
}
