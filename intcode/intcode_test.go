package intcode

import (
	"strconv"
	"strings"
	"testing"
)

func TestInitMemory(t *testing.T) {
	var p = Program{}
	p.InitMemory("1,2,3")
	var expected = []int{1, 2, 3}
	if !slicesEqual(p.memory, expected) {
		t.Errorf("p.memory = %v; want {1, 2, 3}", p.memory)
	}
	if !slicesEqual(p.memoryOrig, expected) {
		t.Errorf("p.memoryOrig = %v; want {1, 2, 3}", p.memoryOrig)
	}
	p.InitMemory("4,5")
	expected = []int{4, 5}
	if !slicesEqual(p.memory, expected) {
		t.Errorf("p.memory = %v; want {4, 5}", p.memory)
	}
	if !slicesEqual(p.memoryOrig, expected) {
		t.Errorf("p.memoryOrig = %v; want {4, 5}", p.memoryOrig)
	}
	p.InitMemory("6,7,8,9")
	expected = []int{6, 7, 8, 9}
	if !slicesEqual(p.memory, expected) {
		t.Errorf("p.memory = %v; want {6, 7, 8, 9}", p.memory)
	}
	if !slicesEqual(p.memoryOrig, expected) {
		t.Errorf("p.memoryOrig = %v; want {6, 7, 8, 9}", p.memoryOrig)
	}
}

func TestRun(t *testing.T) {
	var p = Program{}
	p.InitMemory("1,0,0,0,99")
	p.Run()
	if p.convertMemoryToString() != "2,0,0,0,99" {
		t.Errorf("p.memory = %s; want 2,0,0,0,99", p.convertMemoryToString())
	}

	p.InitMemory("2,3,0,3,99")
	p.Run()
	if p.convertMemoryToString() != "2,3,0,6,99" {
		t.Errorf("p.memory = %s; want 2,3,0,6,99", p.convertMemoryToString())
	}

	p.InitMemory("2,4,4,5,99,0")
	p.Run()
	if p.convertMemoryToString() != "2,4,4,5,99,9801" {
		t.Errorf("p.memory = %s; want 2,4,4,5,99,9801", p.convertMemoryToString())
	}

	p.InitMemory("1,1,1,4,99,5,6,0,99")
	p.Run()
	if p.convertMemoryToString() != "30,1,1,4,2,5,6,0,99" {
		t.Errorf("processInput = %s; want 30,1,1,4,2,5,6,0,99", p.convertMemoryToString())
	}

	p.InitMemory("1,9,10,3,2,3,11,0,99,30,40,50")
	p.Run()
	if p.convertMemoryToString() != "3500,9,10,70,2,3,11,0,99,30,40,50" {
		t.Errorf("processInput = %s; want 3500,9,10,70,2,3,11,0,99,30,40,50", p.convertMemoryToString())
	}
}

func TestAddOp(t *testing.T) {
	var p = Program{}
	p.InitMemory("1,4,4,5,99,0")
	var inst = parseInstruction(p.memory[0])
	var got = addOp(&p, 0, inst)
	if p.convertMemoryToString() != "1,4,4,5,99,198" {
		t.Errorf("addOp = %s; want 1,4,4,5,99,198", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}

	p.InitMemory("11101,4,4,5,99,0")
	inst = parseInstruction(p.memory[0])
	got = addOp(&p, 0, inst)
	if p.convertMemoryToString() != "11101,4,4,8,99,0" {
		t.Errorf("addOp = %s; want 11101,4,4,8,99,0", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}
}

func TestMulOp(t *testing.T) {
	var p = Program{}
	p.InitMemory("2,4,4,5,99,9801")
	var inst = parseInstruction(p.memory[0])
	var got = mulOp(&p, 0, inst)
	if p.convertMemoryToString() != "2,4,4,5,99,9801" {
		t.Errorf("addOp = %s; want 2,4,4,5,99,9801", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}

	p.InitMemory("11102,4,4,5,99,0")
	inst = parseInstruction(p.memory[0])
	got = mulOp(&p, 0, inst)
	if p.convertMemoryToString() != "11102,4,4,16,99,0" {
		t.Errorf("addOp = %s; want 11102,4,4,16,99,0", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}
}

func TestStoreOp(t *testing.T) {
	var p = Program{}
	p.InitMemory("3,3,99,0")
	p.input = []int{15}
	var inst = parseInstruction(p.memory[0])
	var got = storeOp(&p, 0, inst)
	if p.convertMemoryToString() != "3,3,99,15" {
		t.Errorf("addOp = %s; want 3,3,99,15", p.convertMemoryToString())
	}
	if got != 2 {
		t.Errorf("ip = %d; want 2", got)
	}
	if len(p.input) != 0 {
		t.Errorf("p.input = %v; want []", p.input)
	}

	p.InitMemory("103,3,99,0")
	p.input = []int{15}
	inst = parseInstruction(p.memory[0])
	got = storeOp(&p, 0, inst)
	if p.convertMemoryToString() != "103,15,99,0" {
		t.Errorf("addOp = %s; want 103,15,99,0", p.convertMemoryToString())
	}
	if got != 2 {
		t.Errorf("ip = %d; want 2", got)
	}
	if len(p.input) != 0 {
		t.Errorf("p.input = %v; want []", p.input)
	}
}

func TestLoadOp(t *testing.T) {
	var p = Program{}
	p.InitMemory("4,5,4,6,99,20,25")
	var inst = parseInstruction(p.memory[0])
	var got = loadOp(&p, 0, inst)
	got = loadOp(&p, got, inst)
	var expected = []int{20, 25}
	if p.convertMemoryToString() != "4,5,4,6,99,20,25" {
		t.Errorf("loadOp = %s; want 4,5,4,6,99,20,25", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 2", got)
	}
	if !slicesEqual(p.output, expected) {
		t.Errorf("p.output = %v; want [20 25]", p.output)
	}

	p.Reset()
	p.InitMemory("104,5,104,6,99,20,25")
	inst = parseInstruction(p.memory[0])
	got = loadOp(&p, 0, inst)
	got = loadOp(&p, got, inst)
	expected = []int{5, 6}
	if p.convertMemoryToString() != "104,5,104,6,99,20,25" {
		t.Errorf("loadOp = %s; want 104,5,104,6,99,20,25", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 2", got)
	}
	if !slicesEqual(p.output, expected) {
		t.Errorf("p.output = %v; want [5 6]", p.output)
	}
}

func TestJnzOp(t *testing.T) {
	var p = Program{}
	p.InitMemory("5,4,5,99,1,4")
	var inst = parseInstruction(p.memory[0])
	var got = jnzOp(&p, 0, inst)
	if p.convertMemoryToString() != "5,4,5,99,1,4" {
		t.Errorf("p.memory = %s; want 5,4,5,99,1,4", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}

	p.Reset()
	p.InitMemory("5,4,5,99,0,4")
	inst = parseInstruction(p.memory[0])
	got = jnzOp(&p, 0, inst)
	if p.convertMemoryToString() != "5,4,5,99,0,4" {
		t.Errorf("p.memory = %s; want 5,4,5,99,0,4", p.convertMemoryToString())
	}
	if got != 3 {
		t.Errorf("ip = %d; want 3", got)
	}

	p.Reset()
	p.InitMemory("1105,1,4,99,0")
	inst = parseInstruction(p.memory[0])
	got = jnzOp(&p, 0, inst)
	if p.convertMemoryToString() != "1105,1,4,99,0" {
		t.Errorf("p.memory = %s; want 1105,1,4,99,0", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}

	p.Reset()
	p.InitMemory("1105,0,4,99,0")
	inst = parseInstruction(p.memory[0])
	got = jnzOp(&p, 0, inst)
	if p.convertMemoryToString() != "1105,0,4,99,0" {
		t.Errorf("p.memory = %s; want 1105,0,4,99,0", p.convertMemoryToString())
	}
	if got != 3 {
		t.Errorf("ip = %d; want 3", got)
	}
}

func TestJzOp(t *testing.T) {
	var p = Program{}
	p.InitMemory("6,4,5,99,1,4")
	var inst = parseInstruction(p.memory[0])
	var got = jzOp(&p, 0, inst)
	if p.convertMemoryToString() != "6,4,5,99,1,4" {
		t.Errorf("p.memory = %s; want 6,4,5,99,1,4", p.convertMemoryToString())
	}
	if got != 3 {
		t.Errorf("ip = %d; want 3", got)
	}

	p.Reset()
	p.InitMemory("6,4,5,99,0,4")
	inst = parseInstruction(p.memory[0])
	got = jzOp(&p, 0, inst)
	if p.convertMemoryToString() != "6,4,5,99,0,4" {
		t.Errorf("p.memory = %s; want 6,4,5,99,0,4", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}

	p.Reset()
	p.InitMemory("1106,1,4,99,0")
	inst = parseInstruction(p.memory[0])
	got = jzOp(&p, 0, inst)
	if p.convertMemoryToString() != "1106,1,4,99,0" {
		t.Errorf("p.memory = %s; want 1106,1,4,99,0", p.convertMemoryToString())
	}
	if got != 3 {
		t.Errorf("ip = %d; want 3", got)
	}

	p.Reset()
	p.InitMemory("1106,0,4,99,0")
	inst = parseInstruction(p.memory[0])
	got = jzOp(&p, 0, inst)
	if p.convertMemoryToString() != "1106,0,4,99,0" {
		t.Errorf("p.memory = %s; want 1106,0,4,99,0", p.convertMemoryToString())
	}
	if got != 4 {
		t.Errorf("ip = %d; want 4", got)
	}
}

func TestParseInstruction(t *testing.T) {
	var got = parseInstruction(1)
	if got.opcode != 1 {
		t.Errorf("instruction.opcode = %d; want 1", got.opcode)
	}

	got = parseInstruction(99)
	if got.opcode != 99 {
		t.Errorf("instruction.opcode = %d; want 99", got.opcode)
	}

	got = parseInstruction(1099)
	if got.opcode != 99 {
		t.Errorf("instruction.opcode = %d; want 99", got.opcode)
	}

	got = parseInstruction(1001)
	expected := []int{0, 1, 0}
	if !slicesEqual(got.modeMask, expected) {
		t.Errorf("instruction.modeMask = %v; want {0, 1, 0}", got.modeMask)
	}

	got = parseInstruction(10001)
	expected = []int{0, 0, 1}
	if !slicesEqual(got.modeMask, expected) {
		t.Errorf("instruction.modeMask = %v; want {0, 0, 1}", got.modeMask)
	}

	got = parseInstruction(3)
	expected = []int{0}
	if !slicesEqual(got.modeMask, expected) {
		t.Errorf("instruction.modeMask = %v; want {0}", got.modeMask)
	}
}

func TestReset(t *testing.T) {
	var p = Program{}
	p.InitMemory("1,9,10,3,2,3,11,0,99,30,40,50")
	p.Run()
	p.Reset()
	if p.convertMemoryToString() != "1,9,10,3,2,3,11,0,99,30,40,50" {
		t.Errorf("processInput = %s; want 1,9,10,3,2,3,11,0,99,30,40,50", p.convertMemoryToString())
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (p Program) convertMemoryToString() string {
	var strArray []string

	for _, intItem := range p.memory {
		strItem := strconv.Itoa(intItem)

		strArray = append(strArray, strItem)
	}

	return strings.Join(strArray, ",")
}
