package intcode

import (
	"log"
	"math"
	"strconv"
	"strings"
)

/*
Program is a structure for running intcode compiler
*/
type Program struct {
	memoryOrig []int64
	memory     []int64
	input      []int64
	output     []int64
	ip         int
	relBase    int
	isRunning  bool
	isPaused   bool
}

type instruction struct {
	opcode   int
	modeMask []int
}

var instructionLength = map[int]int{
	1:  3,
	2:  3,
	3:  1,
	4:  1,
	5:  2,
	6:  2,
	7:  3,
	8:  3,
	9:  1,
	99: 0,
}

/*
InitMemory parses a string and initializes memory of a program.
This will reset initial memory of an intcode program
*/
func (p *Program) InitMemory(input string) {
	strArray := strings.Split(input, ",")
	var intArray []int64

	for _, strItem := range strArray {
		intItem, err := strconv.ParseInt(strItem, 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		intArray = append(intArray, intItem)
	}

	p.memoryOrig = make([]int64, len(intArray))
	p.memory = make([]int64, len(intArray))
	copy(p.memoryOrig, intArray)
	copy(p.memory, intArray)
	p.ip = 0
	p.relBase = 0
	p.input = []int64{}
	p.output = []int64{}
	p.isRunning = true
	p.isPaused = false
}

/*
CopyMemory initialized program with existing copy of memory passed in as array
*/
func (p *Program) CopyMemory(initMemory []int64) {
	p.memoryOrig = make([]int64, len(initMemory))
	copy(p.memoryOrig, initMemory)
	p.Reset()
}

/*
GetMemory returns current memory of program
*/
func (p *Program) GetMemory() []int64 {
	return p.memory
}

/*
Run perfroms main computation of a program
*/
func (p *Program) Run() {
	var inst instruction

	for p.isRunning && !p.isPaused {
		inst = parseInstruction(int(p.memory[p.ip]))

		switch inst.opcode {
		case 99:
			p.isRunning = false
		case 1:
			p.ip = addOp(p, p.ip, inst)
		case 2:
			p.ip = mulOp(p, p.ip, inst)
		case 3:
			p.ip = storeOp(p, p.ip, inst)
		case 4:
			p.ip = loadOp(p, p.ip, inst)
		case 5:
			p.ip = jnzOp(p, p.ip, inst)
		case 6:
			p.ip = jzOp(p, p.ip, inst)
		case 7:
			p.ip = ltOp(p, p.ip, inst)
		case 8:
			p.ip = eqOp(p, p.ip, inst)
		case 9:
			p.ip = addRelBaseOp(p, p.ip, inst)
		}
	}
}

/*
IsRunning returns true if the program is still executing
*/
func (p *Program) IsRunning() bool {
	return p.isRunning
}

/*
Continue resumes execution of program
*/
func (p *Program) Continue() {
	p.isPaused = false
	p.Run()
}

/*
parseInstruction separates instruction into opcode and array of mode values
*/
func parseInstruction(instValue int) instruction {
	var instRet = instruction{}
	if instValue <= 99 {
		instRet.opcode = instValue
	} else {
		instRet.opcode = int(math.Mod(float64(instValue), 100))
	}

	instRet.modeMask = make([]int, instructionLength[instRet.opcode])
	var maskValue = int((instValue - instRet.opcode) / 100)
	for i := 0; i < instructionLength[instRet.opcode]; i++ {
		if maskValue == 0 {
			instRet.modeMask[i] = 0
		} else {
			instRet.modeMask[i] = int(math.Mod(float64(maskValue), 10))
			maskValue = int((maskValue - instRet.modeMask[i]) / 10)
		}
	}

	return instRet
}

func addOp(p *Program, ip int, inst instruction) int {
	par1Addr := getAddrByMode(p, ip+1, inst.modeMask[0])
	par2Addr := getAddrByMode(p, ip+2, inst.modeMask[1])
	outAddr := getAddrByMode(p, ip+3, inst.modeMask[2])

	p.memory[outAddr] = p.memory[par1Addr] + p.memory[par2Addr]
	return ip + instructionLength[inst.opcode] + 1
}

func mulOp(p *Program, ip int, inst instruction) int {
	par1Addr := getAddrByMode(p, ip+1, inst.modeMask[0])
	par2Addr := getAddrByMode(p, ip+2, inst.modeMask[1])
	outAddr := getAddrByMode(p, ip+3, inst.modeMask[2])

	p.memory[outAddr] = p.memory[par1Addr] * p.memory[par2Addr]
	return ip + instructionLength[inst.opcode] + 1
}

func storeOp(p *Program, ip int, inst instruction) int {

	if len(p.input) == 0 {
		p.isPaused = true
		return ip
	}

	storeAddr := getAddrByMode(p, ip+1, inst.modeMask[0])

	p.memory[storeAddr] = p.input[0]
	p.input = p.input[1:]

	return ip + 2
}

func loadOp(p *Program, ip int, inst instruction) int {
	loadAddr := getAddrByMode(p, ip+1, inst.modeMask[0])

	p.output = append(p.output, p.memory[loadAddr])

	return ip + 2
}

func jnzOp(p *Program, ip int, inst instruction) int {
	cmpAddr := getAddrByMode(p, ip+1, inst.modeMask[0])
	jmpAddr := getAddrByMode(p, ip+2, inst.modeMask[1])

	if p.memory[cmpAddr] != 0 {
		return int(p.memory[jmpAddr])
	}
	return ip + 3
}

func jzOp(p *Program, ip int, inst instruction) int {
	cmpAddr := getAddrByMode(p, ip+1, inst.modeMask[0])
	jmpAddr := getAddrByMode(p, ip+2, inst.modeMask[1])

	if p.memory[cmpAddr] == 0 {
		return int(p.memory[jmpAddr])
	}
	return ip + 3
}

func ltOp(p *Program, ip int, inst instruction) int {
	cmpAddr1 := getAddrByMode(p, ip+1, inst.modeMask[0])
	cmpAddr2 := getAddrByMode(p, ip+2, inst.modeMask[1])
	storeAddr := getAddrByMode(p, ip+3, inst.modeMask[2])

	if p.memory[cmpAddr1] < p.memory[cmpAddr2] {
		p.memory[storeAddr] = 1
	} else {
		p.memory[storeAddr] = 0
	}
	return ip + 4
}

func eqOp(p *Program, ip int, inst instruction) int {
	cmpAddr1 := getAddrByMode(p, ip+1, inst.modeMask[0])
	cmpAddr2 := getAddrByMode(p, ip+2, inst.modeMask[1])
	storeAddr := getAddrByMode(p, ip+3, inst.modeMask[2])

	if p.memory[cmpAddr1] == p.memory[cmpAddr2] {
		p.memory[storeAddr] = 1
	} else {
		p.memory[storeAddr] = 0
	}
	return ip + 4
}

func addRelBaseOp(p *Program, ip int, inst instruction) int {
	relBaseAddr := getAddrByMode(p, ip+1, inst.modeMask[0])

	p.relBase += int(p.memory[relBaseAddr])

	return ip + 2
}

func getAddrByMode(p *Program, ip int, modeMask int) int {
	if modeMask == 1 {
		return ip
	} else if modeMask == 2 {
		retAddr := p.relBase + int(p.memory[ip])
		p.expandMemory(retAddr + 1)
		return retAddr
	}
	retAddr := int(p.memory[ip])
	p.expandMemory(retAddr + 1)
	return retAddr
}

func (p *Program) expandMemory(newSize int) {
	if len(p.memory) < newSize {
		tmpMemory := make([]int64, newSize)
		copy(tmpMemory, p.memory)
		p.memory = tmpMemory
	}
}

/*
SetInput sets input array of a program
*/
func (p *Program) SetInput(input []int64) {
	p.input = input
}

/*
PushInput pushed additional input on to the end of input array
*/
func (p *Program) PushInput(inputVal int64) {
	p.input = append(p.input, inputVal)
}

/*
GetOutput returns output array of a program
*/
func (p Program) GetOutput() []int64 {
	return p.output
}

/*
PopOutput pops output value from front of output array
*/
func (p *Program) PopOutput() int64 {
	var outVal = p.output[0]
	p.output = p.output[1:]
	return outVal
}

/*
SetNounVerb sets noun and verb of a programs memory
*/
func (p *Program) SetNounVerb(noun int64, verb int64) {
	p.memory[1] = noun
	p.memory[2] = verb
}

/*
GetOutputRegister returns output value of a program based on first entry in memory
*/
func (p Program) GetOutputRegister() int64 {
	return p.memory[0]
}

/*
SetMemoryValue will update selected memory location to a new value
*/
func (p *Program) SetMemoryValue(position int, value int64) {
	p.memory[position] = value
}

/*
Reset resets memory to original state
*/
func (p *Program) Reset() {
	p.memory = make([]int64, len(p.memoryOrig))
	copy(p.memory, p.memoryOrig)
	p.input = []int64{}
	p.output = []int64{}
	p.ip = 0
	p.relBase = 0
	p.isRunning = true
	p.isPaused = false
}
