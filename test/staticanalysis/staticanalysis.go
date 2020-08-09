package main

// https://medium.com/a-journey-with-go/go-introduction-to-the-escape-analysis-f7610174e890
// go run -gcflags="-m" staticanalysis.go
import "math/rand"

func main() {
	num := getRandom()
	println(*num)
}

//go:noinline
func getRandom() *int {
	tmp := rand.Intn(100)

	return &tmp
}
