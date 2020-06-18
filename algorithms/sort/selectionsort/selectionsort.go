package main

import "fmt"

func selectionSort(array []int) {

	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		array[i], array[min] = array[min], array[i]
	}
}

func main() {

	arr := []int{9, 1, 2, 8, 6, 7, 5, 3, 0, 4}
	fmt.Println("Initial array is:", arr)
	fmt.Println("")

	/*var min int

	for i := 0; i < len(arr)-1; i++ {
		min = i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	} */

	selectionSort(arr)

	fmt.Println("Sorted array:    ", arr)
}
