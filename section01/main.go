package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	fmt.Println("Hello", os.Args[1])
}
