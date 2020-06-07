package main

import (
	"flag"
	"fmt"
	"github.com/kidmam/TIL/algorithms/spigot"
)

func main() {
	N := flag.Int("n", 60, "number of digits")
	flag.Parse()
	fmt.Println(spigot.Pi(*N))
}
