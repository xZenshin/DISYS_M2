package main

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

	f, err := os.OpenFile("../Token Ring Log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter port number (You can only choose between 5000, 5001 and 5002): ")
	inputPort, _ := reader.ReadString('\n')
	inputPort = strings.TrimRight(inputPort, "\r\n")

	file, err := os.Open("../ports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		if split[0] == inputPort {
			log.Println("Found port")
			id, err := strconv.Atoi(split[2])
			if err != nil {
				fmt.Println("Big error", err)
			}
			node = n.Node{
				Port:         split[0],
				NextNodePort: split[1],
				ID:           id,
			}
			log.Printf("Starting node with port: %s\tNext nodeport: %s\tID: %d\n", node.Port, node.NextNodePort, node.ID)
			break
		}
	}

	if node.ID != 0 {
		go n.ListenForMessages(node)
		go node.TryToAccessCriticalSection()
		if node.Port == "5002" {
			node.NodeStart(node.NextNodePort)
		}
		for {

		}
	} else {
		fmt.Println("Wrong port number")
		main()
	}

}
