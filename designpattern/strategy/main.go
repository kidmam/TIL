package main

import "fmt"

// Be able to override a struct
type A struct {
	year int
}

func (a A) Greet() { fmt.Println("Hello GolangUK", a.year) }

type B struct {
	A
}

func (b B) Greet() { fmt.Println("Welcome to GolangUK", b.year) }

// Be able to override a struct

type Calculer interface {
	Execute(int, int) int
}
type Add struct {
}

func (add Add) Execute(a int, b int) int {
	return a + b
}

type Minus struct {
}

func (m Minus) Execute(a int, b int) int {
	return a - b
}

type Calcul struct {
	c Calculer
}

func (c Calcul) Calculate(a int, b int) int {
	return c.c.Execute(a, b)
}
func main() {
	a := Calcul{c: Add{}}
	b := Calcul{c: Minus{}}
	fmt.Println(a.Calculate(3, 1)) // 4
	fmt.Println(b.Calculate(3, 1)) // 2
}
