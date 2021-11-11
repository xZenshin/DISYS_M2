package Node

import (
	"context"
	"fmt"
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
	fmt.Println("RECIEVED REQUEST FROM", token.IdFrom, " SENT TO", token.PortTo)
	fmt.Println("Current Node ID:", s.ID, " Port", s.Port, " NextPort", s.NextNodePort)
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

func AccessCriticalSection(id int) {
	fmt.Println("ACCESS CRITICAL SECTION: ID ", id)
	time.Sleep(time.Second * 1)
}

func (n *Node) ClientStart(nextPort string) {
	AccessCriticalSection(n.ID)

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
