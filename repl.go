package main

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "bf >> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		// Create compiler
		c := NewCompiler(line)

		// Compile Brainfuck code
		instructions := c.Compile()

		// Create machine
		m := NewMachine(
			instructions,
			in,
			out,
		)

		// Execute instructions
		m.Execute()

		fmt.Fprintln(out)
	}
}
