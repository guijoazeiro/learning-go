package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	//Integer
	var integerVar int = 42
	fmt.Println("int: ", integerVar)

	//Float
	var floatVar float64 = 3.14
	fmt.Println("float: ", floatVar)

	//String
	var stringVar string = "Hello, World!"
	fmt.Println("string: ", stringVar)

	//boolean
	var boolVar bool = true
	fmt.Println("bool: ", boolVar)

	//Array
	var intArray [3]int = [3]int{1, 2, 3}
	fmt.Println("array: ", intArray)

	//Pointer
	var originalVar int = 42
	var pointerVar *int = &originalVar
	fmt.Println("endereço do ponteiro: ", pointerVar)
	fmt.Println("valor do ponteiro: ", *pointerVar)
	*pointerVar = 43
	fmt.Println("valor da variável de origem após alteração no ponteiro: ", originalVar)

	//Struct
	var personVar Person = Person{Name: "Alice", Age: 42}
	fmt.Println("Struct: ", personVar)
	fmt.Println("Valor de um atributo: ", personVar.Name)

	//Map
	var myMap map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Map: ", myMap)

	//Interface
	var myShape Shape = Circle{Radius: 5}
	fmt.Println("Interface: ", myShape.Area())
}
