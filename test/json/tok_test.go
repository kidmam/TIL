package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func BenchmarkJSONDecodeHello(b *testing.B) {
	input := `"hello"`
	r := strings.NewReader(input) // reads strings as []byte
	dec := json.NewDecoder(r)
	b.ReportAllocs()
	b.SetBytes(int64(len(input)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Seek(0, 0)
		tok, _ := dec.Token()
		if tok != "hello" {
			b.Fatal()
		}
	}
}
