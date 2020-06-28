package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	input := `{"a": 1, "b": true, "c": [1, "two", null]}`
	dec := json.NewDecoder(strings.NewReader(input))
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		fmt.Printf("%v\t(%T)\n", tok, tok)
	}

}
