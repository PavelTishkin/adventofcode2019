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
	memoryOrig []int
	memory     []int
	input      []int
	output     []int
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
	99: 0,
}

/*
InitMemory parses a string and initializes memory of a program.
This will reset initial memory of an intcode program
*/
func (p *Program) InitMemory(input string) {
	strArray := strings.Split(input, ",")
	var intArray []int

	for _, strItem := range strArray {
		intItem, err := strconv.ParseInt(strItem, 10, 32)

		if err != nil {
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
Run perfroms main computation of a program
*/
func (p *Program) Run() {
	var isRunning bool = true
	var ip int = 0
	var inst instruction

	for isRunning {
		inst = parseInstruction(p.memory[ip])

		switch inst.opcode {
		case 99:
			isRunning = false
		case 1:
			ip = addOp(p, ip, inst)
		case 2:
			ip = mulOp(p, ip, inst)
		case 3:
			ip = storeOp(p, ip, inst)
		case 4:
			ip = loadOp(p, ip, inst)
		case 5:
			ip = jnzOp(p, ip, inst)
		case 6:
			ip = jzOp(p, ip, inst)
		case 7:
			ip = ltOp(p, ip, inst)
		case 8:
			ip = eqOp(p, ip, inst)
		}
	}
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
	var par1Addr, par2Addr, outAddr int
	if inst.modeMask[0] == 1 {
		par1Addr = ip + 1
	} else {
		par1Addr = p.memory[ip+1]
	}
	if inst.modeMask[1] == 1 {
		par2Addr = ip + 2
	} else {
		par2Addr = p.memory[ip+2]
	}
	if inst.modeMask[2] == 1 {
		outAddr = ip + 3
	} else {
		outAddr = p.memory[ip+3]
	}

	p.memory[outAddr] = p.memory[par1Addr] + p.memory[par2Addr]
	return ip + instructionLength[inst.opcode] + 1
}

func mulOp(p *Program, ip int, inst instruction) int {
	var par1Addr, par2Addr, outAddr int
	if inst.modeMask[0] == 1 {
		par1Addr = ip + 1
	} else {
		par1Addr = p.memory[ip+1]
	}
	if inst.modeMask[1] == 1 {
		par2Addr = ip + 2
	} else {
		par2Addr = p.memory[ip+2]
	}
	if inst.modeMask[2] == 1 {
		outAddr = ip + 3
	} else {
		outAddr = p.memory[ip+3]
	}

	p.memory[outAddr] = p.memory[par1Addr] * p.memory[par2Addr]
	return ip + instructionLength[inst.opcode] + 1
}

func storeOp(p *Program, ip int, inst instruction) int {
	var storeAddr int
	if inst.modeMask[0] == 1 {
		storeAddr = ip + 1
	} else {
		storeAddr = p.memory[ip+1]
	}

	p.memory[storeAddr] = p.input[0]
	p.input = p.input[1:]

	return ip + 2
}

func loadOp(p *Program, ip int, inst instruction) int {
	var loadAddr int
	if inst.modeMask[0] == 1 {
		loadAddr = ip + 1
	} else {
		loadAddr = p.memory[ip+1]
	}

	p.output = append(p.output, p.memory[loadAddr])

	return ip + 2
}

func jnzOp(p *Program, ip int, inst instruction) int {
	var cmpAddr, jmpAddr int
	if inst.modeMask[0] == 1 {
		cmpAddr = ip + 1
	} else {
		cmpAddr = p.memory[ip+1]
	}
	if inst.modeMask[1] == 1 {
		jmpAddr = ip + 2
	} else {
		jmpAddr = p.memory[ip+2]
	}

	if p.memory[cmpAddr] != 0 {
		return p.memory[jmpAddr]
	}
	return ip + 3
}

func jzOp(p *Program, ip int, inst instruction) int {
	var cmpAddr, jmpAddr int
	if inst.modeMask[0] == 1 {
		cmpAddr = ip + 1
	} else {
		cmpAddr = p.memory[ip+1]
	}
	if inst.modeMask[1] == 1 {
		jmpAddr = ip + 2
	} else {
		jmpAddr = p.memory[ip+2]
	}

	if p.memory[cmpAddr] == 0 {
		return p.memory[jmpAddr]
	}
	return ip + 3
}

func ltOp(p *Program, ip int, inst instruction) int {
	var cmpAddr1, cmpAddr2, storeAddr int
	if inst.modeMask[0] == 1 {
		cmpAddr1 = ip + 1
	} else {
		cmpAddr1 = p.memory[ip+1]
	}
	if inst.modeMask[1] == 1 {
		cmpAddr2 = ip + 2
	} else {
		cmpAddr2 = p.memory[ip+2]
	}
	if inst.modeMask[2] == 1 {
		storeAddr = ip + 3
	} else {
		storeAddr = p.memory[ip+3]
	}

	if p.memory[cmpAddr1] < p.memory[cmpAddr2] {
		p.memory[storeAddr] = 1
	} else {
		p.memory[storeAddr] = 0
	}
	return ip + 4
}

func eqOp(p *Program, ip int, inst instruction) int {
	var cmpAddr1, cmpAddr2, storeAddr int
	if inst.modeMask[0] == 1 {
		cmpAddr1 = ip + 1
	} else {
		cmpAddr1 = p.memory[ip+1]
	}
	if inst.modeMask[1] == 1 {
		cmpAddr2 = ip + 2
	} else {
		cmpAddr2 = p.memory[ip+2]
	}
	if inst.modeMask[2] == 1 {
		storeAddr = ip + 3
	} else {
		storeAddr = p.memory[ip+3]
	}

	if p.memory[cmpAddr1] == p.memory[cmpAddr2] {
		p.memory[storeAddr] = 1
	} else {
		p.memory[storeAddr] = 0
	}
	return ip + 4
}

/*
SetInput sets input array of a program
*/
func (p *Program) SetInput(input []int) {
	p.input = input
}

/*
PushInput pushed additional input on to the end of input array
*/
func (p *Program) PushInput(inputVal int) {
	p.input = append(p.input, inputVal)
}

/*
GetOutput returns output array of a program
*/
func (p Program) GetOutput() []int {
	return p.output
}

/*
PopOutput pops output value from front of output array
*/
func (p *Program) PopOutput() int {
	var outVal = p.output[0]
	p.output = p.output[1:]
	return outVal
}

/*
SetNounVerb sets noun and verb of a programs memory
*/
func (p *Program) SetNounVerb(noun int, verb int) {
	p.memory[1] = noun
	p.memory[2] = verb
}

/*
GetOutputRegister returns output value of a program based on first entry in memory
*/
func (p Program) GetOutputRegister() int {
	return p.memory[0]
}

/*
Reset resets memory to original state
*/
func (p *Program) Reset() {
	p.memory = make([]int, len(p.memoryOrig))
	copy(p.memory, p.memoryOrig)
	p.input = []int{}
	p.output = []int{}
}
