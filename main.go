package main

// Go init name is "tokenring"
import (
	"bufio"
	"fmt"
	"log"
	"os"
	n "tokenring/Node"
)

var Ports []string
var Nodes []n.Node

func main() {

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

	setupNodeServer()

	fmt.Printf("nodes length: %d", len(Nodes))

	// Iterate through all nodes and switch between each node being in the Critical Section

}

func setupNodeServer() {

	for i := 0; i < len(Ports); i++ {

		var nextID = i

		if i == len(Ports)-1 {
			nextID = 0
		}

		node := n.Node{
			ID:           i,
			Port:         Ports[i],
			NextNodePort: Ports[nextID],
		}
		Nodes = append(Nodes, node)
	}

	for _, node := range Nodes {
		go node.CreateServer()
		fmt.Println(node.Port)
	}

	//Run forever to let go routines run
	for {
	}
}

func grantNodeAccess(nodeID int) {

	// Grant the node with id 'nodeID' access

}
