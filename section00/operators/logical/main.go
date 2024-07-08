package main

import "fmt"

func main() {
	//AND (E)
	x := true
	y := false
	result := x && y
	fmt.Println("AND: ", result)

	//OR (OU)
	a := true
	b := false
	resultOr := a || b
	fmt.Println("OR: ", resultOr)

	//NOT (NÃO)
	isSunny := true
	isRainy := !isSunny
	fmt.Println("NOT: ", isRainy)
}
