package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type mySlice []int

func (ms mySlice) Len() int {
	return len(ms)
}

func (ms mySlice) Less(i, j int) bool {
	return ms[i] < ms[j]
}

func (ms mySlice) Swap(i, j int) {
	temp := ms[i]
	ms[i] = ms[j]
	ms[j] = temp
}

func main() {
	ms := mySlice{}
	for i := 0; i < 10; i++ {
		ms = append(ms, rand.Intn(100))
	}
	fmt.Println("pre-sort:", ms)
	sort.Sort(ms)
	fmt.Println("post-sort:", ms)
}

