package main

import "fmt"

func main() {
	fmt.Print(LinearSearch([]int{2, 3, 6, 8, 10}, 10))
}

func LinearSearch(data []int, searchVal int) bool {
	for _, key := range data {
		if key == searchVal {
			return true
		}
	}
	return false
}
