package intcode

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
Program - structure for running intcode compiler
*/
type Program struct {
	memoryOrig []int
	memory     []int
}

/*
Parses a string and initializes memory of a program.
This will reset initial memory of an intcode program
*/
func (p *Program) InitMemory(input string) {
	strArray := strings.Split(input, ",")
	var intArray []int

	for _, strItem := range strArray {
		intItem, err := strconv.ParseInt(strItem, 10, 32)

		if err != nil {
			fmt.Println(strItem)
			log.Fatal(err)
		}

		intArray = append(intArray, int(intItem))
	}

	p.memoryOrig = make([]int, len(intArray))
	p.memory = make([]int, len(intArray))
	copy(p.memoryOrig, intArray)
	copy(p.memory, intArray)
}

/*
Perfroms main computation of a program
*/
func (p *Program) Run() {
	var isRunning bool = true
	var ip int = 0
	var nextOp, noun, verb, sp int

	for isRunning {
		nextOp = p.memory[ip]

		switch nextOp {
		case 99:
			isRunning = false
		case 1:
			noun = p.memory[ip+1]
			verb = p.memory[ip+2]
			sp = p.memory[ip+3]

			p.memory[sp] = p.memory[noun] + p.memory[verb]
		case 2:
			noun = p.memory[ip+1]
			verb = p.memory[ip+2]
			sp = p.memory[ip+3]

			p.memory[sp] = p.memory[noun] * p.memory[verb]
		}

		ip += 4
	}
}

/*
Set input parameters to program
*/
func (p *Program) SetInput(noun int, verb int) {
	p.memory[1] = noun
	p.memory[2] = verb
}

/*
Return output value of a program
*/
func (p Program) GetOutput() int {
	return p.memory[0]
}

/*
Reset memory to original state
*/
func (p *Program) Reset() {
	p.memory = make([]int, len(p.memoryOrig))
	copy(p.memory, p.memoryOrig)
}
