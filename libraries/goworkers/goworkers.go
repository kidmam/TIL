package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dpaks/goworkers"
)

func main() {
	opts := goworkers.Options{Workers: 20}
	gw := goworkers.New(opts)

	// your actual work
	fn := func(i int) {
		fmt.Println("Start Job", i)
		time.Sleep(time.Duration(i) * time.Second)
		fmt.Println("End Job", i)
	}

	for _, value := range []int{9, 7, 1, 2, 3} {
		i := value
		gw.Submit(func() {
			fn(i)
		})
	}
	log.Println("Submitted!")

	gw.Stop()
}
