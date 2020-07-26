package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{4, 1, 9, 3, 8}

	sort.Ints(ints)

	fmt.Println(ints)
	// [1, 3, 4, 8, 9]
}
