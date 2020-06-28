package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := `"hello"`
	var r io.Reader = strings.NewReader(input) // reads strings as []byte
	dec := json.NewDecoder(r)
	tok, _ := dec.Token()
	fmt.Printf("%s\t(%T)\n", tok, tok)
}
