package main

// https://bitfieldconsulting.com/golang/map-declaring-initializing

import "fmt"

func main() {
	var menu map[string]float64
	menu = map[string]float64{
		"eggs":    1.75,
		"bacon":   3.22,
		"sausage": 1.89,
	}

	fmt.Printf("menu=", menu)
}
