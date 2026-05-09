package main

type InsType byte

const (
	Plus          InsType = '+'
	Minus         InsType = '-'
	Right         InsType = '>'
	Left          InsType = '<'
	PutChar       InsType = '.'
	ReadChar      InsType = ','
	JumpIfZero    InsType = '['
	JumpIfNotZero InsType = ']'
)

// Example uses:
// Use 1: +++++ will be turned into Instruction{Type: Plus, Argument: 5}
// Use 2: will contain the position of the instruction of the matching bracket
// For example, for the code: [++], the instruction for the first [ will be Instruction{Type: JumpIfZero, Argument: 3} and the instruction for the second ] will be Instruction{Type: JumpIfNotZero, Argument: 0}
type Instruction struct {
	Type     InsType
	Argument int
}
