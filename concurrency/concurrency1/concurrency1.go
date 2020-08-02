package main

// https://itnext.io/comparing-crystals-concurrency-with-that-of-go-s-part-i-cd45a3388935
import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
