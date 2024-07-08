package main

import "fmt"

func main() {
	var singleVariable int
	singleVariable = 10
	fmt.Println(singleVariable)

	//criação combinada de variaveis
	var a, b, c = 1, 2, "hello"
	fmt.Println(a, b, c)

	//create and assign
	createAndAssignVar := 42
	fmt.Println(createAndAssignVar)

	//using var
	var variable1 int = 5
	variable1 = 10 // Valid, the value of variable1 can be changed

	// using const
	const constant1 int = 5
	//constant1 = 10 //invalid, constants cannot be reassigned

	fmt.Println(variable1, constant1)
}
