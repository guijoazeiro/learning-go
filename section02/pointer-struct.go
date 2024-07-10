package main

import "fmt"

type Empregado struct {
	Nome string
	id   int
}

func main() {
	emp := Empregado{"Glaucio", 123}
	pts := &emp

	fmt.Println(pts)
	pts.Nome = "Guerra"
	fmt.Println("valor pts ", pts)
}
