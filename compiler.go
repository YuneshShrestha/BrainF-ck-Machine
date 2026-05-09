package main

// Eg: For Code: +++++
// code = "+++++"
// codeLength = 5
// position = 0
// instructions = []{Instruction{Type: Plus, Argument: 5}}
type Compiler struct {
	code         string         // Original Brainfuck code
	codeLength   int            // total characters
	position     int            // Current position in code
	instructions []*Instruction // Compiled optimized instructions
}

func NewCompiler(code string) *Compiler {
	return &Compiler{code: code, codeLength: len(code), instructions: []*Instruction{}}
}

func (c *Compiler) Compile() []*Instruction {
	loopStack := []int{}
	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '+':
			c.CompileFoldableInstruction('+', Plus)
		case '-':
			c.CompileFoldableInstruction('-', Minus)
		case '>':
			c.CompileFoldableInstruction('>', Right)
		case '<':
			c.CompileFoldableInstruction('<', Left)
		case '.':
			c.CompileFoldableInstruction('.', PutChar)
		case ',':
			c.CompileFoldableInstruction(',', ReadChar)
		case '[':
			// Initally set the argument to 0 later it will be updated in "]" case
			insPos := c.EmitWithArg(JumpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			// Get the last "[" position from the loopStack
			openInstruction := loopStack[len(loopStack)-1]
			// Remove the last "[" position from the loopStack
			loopStack = loopStack[:len(loopStack)-1]
			// Emit the close instruction "]" with the open instruction position
			closeInstructionPos := c.EmitWithArg(JumpIfNotZero, openInstruction)
			// Update the argument of the open instruction with the close instruction position
			c.instructions[openInstruction].Argument = closeInstructionPos
		}

		c.position++
	}
	return c.instructions
}

func (c *Compiler) CompileFoldableInstruction(char byte, insType InsType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}
	c.EmitWithArg(insType, count)
}

func (c *Compiler) EmitWithArg(insType InsType, arg int) int {
	ins := &Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}
