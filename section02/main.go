package main

import "fmt"

func main() {
	var inteiro = 45
	var ponteiro *int = &inteiro

	fmt.Println("Valor do inteiro: ", inteiro)
	fmt.Println("Endereço do inteiro: ", &inteiro)

	fmt.Println("Endereço do ponteiro: ", ponteiro)
	fmt.Println("Valor do ponteiro: ", *ponteiro)

	inteiro = 55
	fmt.Println("Novo valor do inteiro: ", inteiro)

	*ponteiro = 65
	fmt.Println("Novo valor do ponteiro: ", *ponteiro)
	fmt.Println("Novo valor do inteiro: ", inteiro)
}
