package main

import "fmt"

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
