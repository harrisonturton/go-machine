package instruction

import "fmt"

type Instruction int

const (
	PSH Instruction = iota
	ADD
	POP
	BR
	HLT
)

var instructionDisplay = [...]string{
	"PSH",
	"ADD",
	"POP",
	"BR",
	"HLT",
}

// Convert the instruction to a string
func (i Instruction) String() string {
	if 0 <= i && i < Instruction(len(instructionDisplay)) {
		return instructionDisplay[i]
	} 

	return fmt.Sprintf("Instruction(%s)", string(int(i)))
}
