package main

func main() {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}:
		// Panic if the first case is selected.
	case <-c:
	}
}
