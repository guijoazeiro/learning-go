package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("erro, não é um número valido")
		os.Exit(1)
	}

	fmt.Println("numero convertido", n)

	/*pessoa := &funcionarios.Pessoa{
		nome:    "John",
		idade:   25,
		salario: 100,
	}

	salFuncs := make(map[string]int)
	salFuncs["John"] = 100
	salFuncs["Jane"] = 200

	sal, exists := salFuncs["John"]
	fmt.Println(sal, exists)
	totalSal := len(salFuncs)
	fmt.Println(totalSal)
	*/
	// salarios := []int{}

	// // salarios := make([]int, 5)

	// for i := 0; i < 5; i++ {
	// 	salarios = append(salarios, 100+i)
	// }

	// for _, salario := range salarios {
	// 	fmt.Println(salario)
	// }

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

// func setName(name string) {
// 	fmt.Println("Hello", name)
// }

// func addSalary(valorAtual int, bonus int) (int, int) {
// 	return valorAtual + bonus, bonus
// }

// func (p *Pessoa) addSalaryPessoa(bonus int) {
// 	p.salario += bonus
// }

// func getName() string {
// 	return "John Doe"
// }
