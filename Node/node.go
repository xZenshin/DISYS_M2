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
	Token        pb.Token
}

//Starts another Node-"Client" with NextPort info
func (s *Node) GrantToken(ctx context.Context, token *pb.Token) (*pb.Reply, error) {
	log.Printf("Node %d SENT TO NODE %d", token.IdFrom, ConvertPortToId(token.PortTo))
	log.Printf("Current Node ID: %d PORT: %s NEXTPORT: %s", s.ID, s.Port, s.NextNodePort)
	s.Token = *token
	go s.ClientStart(s.NextNodePort)
	return &pb.Reply{Message: "Token Given"}, nil
}

//Listen for incoming messages
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

//Accesses the critical section then dials the next port with GrantToken
func (n *Node) ClientStart(nextPort string) {
	n.TryToAccessCriticalSection()
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+nextPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to port %s - %s", n.NextNodePort, err)
	}
	defer conn.Close()
	c := pb.NewTokenRingClient(conn)

	response, err := c.GrantToken(context.Background(), &pb.Token{Message: "Secret Code", IdFrom: int32(n.ID), PortTo: n.NextNodePort})
	if err != nil {
		log.Fatalf("Error when calling GrantToken: %s", err)
	}
	log.Printf("Response from Node: %s", response.GetMessage())
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

// Permanently try to access the criticalsection
func (n *Node) TryToAccessCriticalSection() {
	if n.Token.Message == "Secret Code" {
		n.AccessCriticalSection()
	}
}

// After accessing change token message such that this Node no longer will be able to access again
func (n *Node) AccessCriticalSection() {
	log.Printf("ACCESS CRITICAL SECTION ID: %d", n.ID)
	n.Token.Message = ""
	time.Sleep(time.Second * 1)
}
