package main

import (
	"fmt"
	"math/bits"
)

// https://yourbasic.org/golang/bitwise-operator-cheat-sheet/
func main() {
	fmt.Print(bits.UintSize)

	fmt.Print(bits.OnesCount8(00101110))
}
