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
