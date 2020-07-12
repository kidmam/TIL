package dfs

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Data  interface{}
	Sleep time.Duration
	Left  *Node
	Right *Node
}

func NewNode(data interface{}) *Node {

	node := new(Node)

	node.Data = data
	node.Left = nil
	node.Right = nil

	rand.Seed(time.Now().UTC().UnixNano())
	duration := int64(rand.Intn(100))
	node.Sleep = time.Duration(duration) * time.Microsecond

	return node
}

func (n *Node) ProcessNode() {

	var hello []int

	for i := 0; i < 10000; i++ {
		time.Sleep(n.Sleep)
		hello = append(hello, i)
	}

	fmt.Printf("Node %v ✅\n", n.Data)
}
