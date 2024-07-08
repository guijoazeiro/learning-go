package main

import "fmt"

func main() {
	// maior que
	x := 8
	y := 5
	isGreater := x > y
	fmt.Println("x maior que y", isGreater)

	// menor que
	p := 12
	q := 18
	isLess := p < q
	fmt.Println("p menor que q", isLess)

	// maior ou igual
	m := 5
	n := 5
	isGreaterOrEqual := m >= n
	fmt.Println("m maior ou igual a n", isGreaterOrEqual)

	// menor ou igual
	r := 10
	s := 15
	isLessOrEqual := r <= s
	fmt.Println("r menor ou igual a s", isLessOrEqual)

	// igual
	age := 25
	checkAge := age == 25
	fmt.Println("age é igual a 25", checkAge)

	// diferente
	score1 := 80
	passingScore := score1 != 75
	fmt.Println("score1 é diferente de 75", passingScore)
}
