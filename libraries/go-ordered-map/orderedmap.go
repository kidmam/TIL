package main

// https://morioh.com/p/990229f32171?f=5c224490c513a556c9042463
import (
	"fmt"

	"github.com/wk8/go-ordered-map"
)

func main() {
	om := orderedmap.New()

	om.Set("foo", "bar")
	om.Set("bar", "baz")
	om.Set("coucou", "toi")

	fmt.Println(om.Get("foo"))          // => bar, true
	fmt.Println(om.Get("i dont exist")) // => <nil>, false

	// iterating pairs from oldest to newest:
	for pair := om.Oldest(); pair != nil; pair = pair.Next() {
		fmt.Printf("%s => %s\n", pair.Key, pair.Value)
	} // prints:
	// foo => bar
	// bar => baz
	// coucou => toi

	// iterating over the 2 newest pairs:
	i := 0
	for pair := om.Newest(); pair != nil; pair = pair.Prev() {
		fmt.Printf("%s => %s\n", pair.Key, pair.Value)
		i++
		if i >= 2 {
			break
		}
	} // prints:
	// coucou => toi
	// bar => baz
}
