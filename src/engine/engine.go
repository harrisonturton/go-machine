package engine

import (
	"fmt"
	"virtual-machine/src/instruction"
)

type Engine struct {
	program   []int
	registers []int
	stack     []int
	pc        int
	lr        int
	sp        int
	isRunning bool
}

// Create a new instance of Engine
func NewEngine(program []int) *Engine {
	return &Engine{
		program:   program,
		registers: make([]int, 13),
		stack:     make([]int, 256),
		pc:        0,
		lr:        0,
		sp:        -1, // Simplifies creation of e.push()
		isRunning: true,
	}
}

func (e *Engine) Run() {
	for e.isRunning {
		fmt.Println(e.stack[:5])
		e.eval(e.fetch())
		e.pc++
	}
}

// Fetch the next instruction, move pc forward
func (e *Engine) fetch() int {
	if e.pc >= len(e.program) {
		return int(instruction.HLT)
	}
	return e.program[e.pc]
}

// Perform the instruction
func (e *Engine) eval(opcode int) {
	switch instruction.Instruction(opcode) {
	case instruction.HLT:
		e.isRunning = false
		fmt.Println(fmt.Sprintf("Halting at instruction %d", e.pc))
	case instruction.POP:
		e.pop()
	case instruction.PSH:
		e.pc++
		e.push(e.fetch())
	case instruction.BR:
		e.pc++
		fmt.Println(fmt.Sprintf("Moving to %d", e.fetch()))
		e.pc = e.fetch() - 1
	case instruction.ADD:
		a := e.pop()
		b := e.pop()
		e.push(a + b)
	}
}

// Push a new value to the stack
func (e *Engine) push(val int) {
	e.sp++
	e.stack[e.sp] = val
}

// Pop a value off the stack
func (e *Engine) pop() int {
	val := e.stack[e.sp]
	e.stack[e.sp] = 0
	e.sp--
	fmt.Println(fmt.Sprintf("Popped %d", val))
	return val
}
