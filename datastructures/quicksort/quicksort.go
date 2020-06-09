package main

import (
	"fmt"
	"sort"
)

func main() {
	intArray := []int{0, 2, 3, 4, 11, 34}
	searchNumber := 34

	sorted := sort.SearchInts(intArray, searchNumber)

	if sorted < len(intArray) {
		fmt.Printf("Found element %d at array position %d \n", searchNumber, sorted)
	} else {
		fmt.Printf("Element %d not found in array %v \n", searchNumber, intArray)
	}

}
