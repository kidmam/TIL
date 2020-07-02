package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func change(i int) {

	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	fmt.Printf("[%d , %d]", v1, i)
	if v1%10 == 0 {
		fmt.Printf("{%d , %d}", v1, i)
		v1 = v1 - 10*i
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {

	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup

	fmt.Printf("%d", read())

	for i := 0; i < num; i++ {
		waitGroup.Add(1)

		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf("->%d", read())
		}(i)
	}

	waitGroup.Wait()
	fmt.Printf("->%d \n", read())
}
