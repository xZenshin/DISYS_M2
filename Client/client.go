package Client

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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter port number (You can only choose between 5000, 5001 and 5002): ")
	inputPort, _ := reader.ReadString('\n')

	file, err := os.Open("ports.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		if split[0] == inputPort {
			id, err := strconv.Atoi(split[2])
			if err != nil {
			}
			node = n.Node{
				Port:         split[0],
				NextNodePort: split[1],
				ID:           id,
			}
			break
		}
	}

	if node.ID != 0 {

		go n.ListenForMessages(node)
		go node.TryToAccessCriticalSection()

		for {

		}
	}

}
