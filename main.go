package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Brainfuck REPL")
	fmt.Println("Type Brainfuck code below")

	// Start REPL
	Start(os.Stdin, os.Stdout)
}
