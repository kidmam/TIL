package main

// https://www.freecodecamp.org/news/supercharge-your-dfs-with-goroutines/

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
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

func (n *Node) ProcessNodeParallel() {
	defer wg.Done()

	var hello []int

	for i := 0; i < 10000; i++ {
		time.Sleep(n.Sleep)
		hello = append(hello, i)
	}

	fmt.Printf("Node %v ✅\n", n.Data)
}

func (n *Node) DFS() {
	if n == nil {
		return
	}

	n.Left.DFS()
	n.ProcessNode()
	n.Right.DFS()
}

func (n *Node) DFSParallel() {
	defer wg.Done()

	if n == nil {
		return
	}

	wg.Add(1)
	go n.Left.DFSParallel()

	wg.Add(1)
	go n.ProcessNodeParallel()

	wg.Add(1)
	go n.Right.DFSParallel()
}

var wg sync.WaitGroup

func main() {
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)

	start := time.Now()
	root.DFS()
	fmt.Printf("\nTime elapsed: %v\n\n", time.Since(start))

	processors := runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("\nTime elapsed: %v\n\n", time.Since(start))

	// Starts the timer
	start = time.Now()

	// Adds one goroutine the WaitGroup
	wg.Add(1)
	// Start the DFS Goroutine
	go root.DFSParallel()
	// Waits for all goroutines to complete
	wg.Wait()

	fmt.Printf("\nProcessors: %v Time elapsed: %v\n", processors, time.Since(start))

}
