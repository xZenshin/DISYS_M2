package main

//go mod init is called Client
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	n "tokenring/Node"
)

var (
	node n.Node
)

func main() {
	/*
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter port number (You can only choose between 5000, 5001 and 5002): ")
	inputPort, _ := reader.ReadString('\n')

	file, err := os.Open("../ports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line 
		split := strings.Split(scanner.Text(), " ")
		fmt.Println(split)
		if split[0] == inputPort {
			id, err := strconv.Atoi(split[2])
			if err != nil {
				fmt.Println("Big error", err)
			}
			node = n.Node{
				Port:         split[0],
				NextNodePort: split[1],
				ID:           id,
			}
			fmt.Printf("Starting node with port: %s\tNext nodeport: %s\tID: %d\n", node.Port, node.NextNodePort, node.ID)
			break
		}
	}
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter port number (You can only choose between 5000, 5001 and 5002): ")
	inputPort, _ := reader.ReadString('\n')

	node.Port = inputPort
	node.NextNodePort = 

	if node.ID != 0 {
		go n.ListenForMessages(node)
		go node.TryToAccessCriticalSection()

		for {

		}
	}

}
