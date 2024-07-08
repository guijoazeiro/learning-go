package main

import (
	"fmt"
)

func main() {

	name, salario := "John", 100
	setName(name)
	newSalary, bonus := addSalary(salario, 10)
	fmt.Println("Novo Salário:", newSalary)
	fmt.Println("Bonus:", bonus)
	getName()

}

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}

func getName() string {
	return "John Doe"
}
