package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"c", "a", "b"}

	sort.Strings(str)

	fmt.Println(str)
	// [a, b, c]
}
