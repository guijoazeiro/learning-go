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

	salarios := []int{}

	// salarios := make([]int, 5)

	for i := 0; i < 5; i++ {
		salarios = append(salarios, 100+i)
	}

	for _, salario := range salarios {
		fmt.Println(salario)
	}

	// pessoa2 := new(Pessoa)
	// pessoa2.nome = "John2"
	// pessoa2.idade = 25
	// pessoa2.salario = 100

	// fmt.Println("Nome:", pessoa2.nome)

	// pessoa := &Pessoa{
	// 	nome:    "John",
	// 	idade:   25,
	// 	salario: 100,
	// }

	// pessoa.addSalaryPessoa(10)
	// fmt.Println("Novo Salário:", pessoa.salario)

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

func (p *Pessoa) addSalaryPessoa(bonus int) {
	p.salario += bonus
}

func getName() string {
	return "John Doe"
}
