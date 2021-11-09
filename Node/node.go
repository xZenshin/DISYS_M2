package Node

import (
	"context"
	"log"
	"net"
	pb "tokenring/DISYS_M2"

	"google.golang.org/grpc"
)

type Node struct {
	ID           int
	Port         string
	NextNodePort string
}

// Server ....
type Server struct {
	pb.UnimplementedTokenRingServer
}

// SayHello ...
func (s *Node) grantToken(ctx context.Context, in *pb.Token) (*pb.Reply, error) {
	log.Printf("Receive message body from client: %s", in.Message)
	return &pb.Reply{Message: "Hello From the Server!"}, nil
}

func (n *Node) ServerStart() {
	for {
		lis, err := net.Listen("tcp", ":"+n.Port)
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}
		s := Server{}
		grpcServer := grpc.NewServer()
		pb.RegisterTokenRingServer(grpcServer, &s)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %s", err)
		}
	}
}

func (n *Node) ClientStart() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":"+n.NextNodePort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect to port %s - %s", n.NextNodePort, err)
	}
	defer conn.Close()
	c := pb.NewTokenRingClient(conn)
	response, err := c.GrantToken(context.Background(), &pb.Token{Message: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling GrantToken: %s", err)
	}
	log.Printf("Response from server: %s", response.GetMessage())
}
