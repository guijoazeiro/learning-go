package main

import (
	"fmt"
)

func main() {

	name := getName()
	idade := 26

	fmt.Println("Hello", name, idade)
}

func getName() string {
	return "John Doe"
}
