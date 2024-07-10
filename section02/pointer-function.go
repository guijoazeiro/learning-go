package main

import (
	"fmt"
)

func pointerFunction(a *int) {
	*a = 400
}

func main() {
	var x = 100
	fmt.Printf("o valor da variavel x é: %d\n", x)
	var pointerAddress *int = &x

	pointerFunction(pointerAddress)

	fmt.Printf("o valor da variavel x depois da chamada da função é: %d\n", x)
}
