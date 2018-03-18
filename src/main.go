package main

import (
	"virtual-machine/src/engine"
	"virtual-machine/src/instruction"
)

var program = []int{
	int(instruction.PSH),
	3,
	int(instruction.PSH),
	6,
	int(instruction.ADD),
	int(instruction.BR),
	8,
	int(instruction.HLT),
	int(instruction.PSH),
	1,
	int(instruction.ADD),
	int(instruction.POP),
	int(instruction.HLT),
}

func main() {
	e := engine.NewEngine(program)
	e.Run()	
}
