package main

import "testing"

func TestF(t *testing.T) {
	v := f()
	if v != 4 {
		t.Error("Fails")
	}
}
