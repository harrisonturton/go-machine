package main

import (
	"fmt"
	"virtual-machine/src/instruction"
)

var program = [...]int{
	int(instruction.PSH), 3,
	int(instruction.PSH), 6,
	int(instruction.ADD),
	int(instruction.POP),
	int(instruction.HLT),
}

var stack = make([]int, 256)

var pc = 0
var sp = -1
var running = true

func fetch() int {
	return program[pc]
}

func eval(instr int) {
	switch instruction.Instruction(instr) {
	case instruction.HLT:
		running = false
	case instruction.PSH:
		sp++
		pc++
		val := fetch()
		stack[sp] = val
	case instruction.POP:
		val := stack[sp]
		stack[sp] = 0
		sp--
		fmt.Println(val)
	case instruction.ADD:
		first := stack[sp]
		stack[sp] = 0
		sp--
		second := stack[sp]
		stack[sp] = 0
		sp--
		sp++
		stack[sp] = first + second
	}
}

func main() {
	for running {
		fmt.Println(stack[:5])
		eval(fetch())
		pc++
	}
}
