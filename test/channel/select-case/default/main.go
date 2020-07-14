package main

import "fmt"

func main() {
	var c chan struct{} // nil
	select {
	case <-c: // blocking operation
	case c <- struct{}{}: // blocking operation
	default:
		fmt.Println("Go here.")
	}
}
