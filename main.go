package main

// Go init name is "tokenring"
import (
	"bufio"
	"fmt"
	"log"
	"os"
	pb "tokenring/DISYS_M2"
	n "tokenring/Node"
)

var Ports []string
var Nodes []n.Node

func main() {
	f, err := os.OpenFile("Token Ring Log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("START OF TOKENRING EXAMPLE")

	file, err := os.Open("ports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		Ports = append(Ports, scanner.Text())
	}
	fmt.Println(Ports)

	StartNodes()

	fmt.Printf("nodes length: %d", len(Nodes))
}

func StartNodes() {

	for i := 0; i < len(Ports); i++ {

		var nextID = i + 1

		if i == len(Ports)-1 {
			nextID = 0
		}

		node := n.Node{
			ID:           i,
			Port:         Ports[i],
			NextNodePort: Ports[nextID],
		}
		Nodes = append(Nodes, node)
		Nodes[0].Token = pb.Token{Message: "Secret Code"}
	}

	for _, node := range Nodes {
		go n.ListenForMessages(node)
		log.Printf("Started server with port: " + node.Port)
	}
	for _, nodee := range Nodes {
		nodee.TryToAccessCriticalSection()
	}

	Nodes[0].NodeStart(Nodes[0].NextNodePort)

	//Run forever to let go routines run
	for {

	}
}
