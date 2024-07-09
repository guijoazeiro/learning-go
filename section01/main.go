package main

import (
	"fmt"
)

type Pessoa struct {
	nome    string
	idade   int
	salario int
}

func main() {

	pessoa := &Pessoa{
		nome:    "John",
		idade:   25,
		salario: 100,
	}

	addSalaryPessoa(pessoa, 10)
	fmt.Println("Novo Salário:", pessoa.salario)

	// name, salario := "John", 100
	// setName(name)
	// newSalary, bonus := addSalary(salario, 10)
	// fmt.Println("Novo Salário:", newSalary)
	// fmt.Println("Bonus:", bonus)
	// getName()

}

func setName(name string) {
	fmt.Println("Hello", name)
}

func addSalary(valorAtual int, bonus int) (int, int) {
	return valorAtual + bonus, bonus
}

func addSalaryPessoa(p *Pessoa, bonus int) {
	p.salario += bonus
}

func getName() string {
	return "John Doe"
}
