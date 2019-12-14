package main

import (
	"testing"

	"../intcode"
)

func TestGetMaxPhaseValue(t *testing.T) {
	p := intcode.Program{}
	p.InitMemory("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")
	phaseArr := []int{0, 1, 2, 3, 4}
	got := getMaxPhaseValue(&p, phaseArr)
	if got != 43210 {
		t.Errorf("getMaxPhaseValue = %d; want 43210", got)
	}

	p.Reset()
	p.InitMemory("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")
	phaseArr = []int{0, 1, 2, 3, 4}
	got = getMaxPhaseValue(&p, phaseArr)
	if got != 54321 {
		t.Errorf("getMaxPhaseValue = %d; want 54321", got)
	}

	p.Reset()
	p.InitMemory("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")
	phaseArr = []int{0, 1, 2, 3, 4}
	got = getMaxPhaseValue(&p, phaseArr)
	if got != 65210 {
		t.Errorf("getMaxPhaseValue = %d; want 65210", got)
	}
}

func TestGetAllPermutations(t *testing.T) {
	initArray := []int{0, 1, 2}
	perms := getAllPermutationsRec(initArray)
	if len(perms) != 6 {
		t.Errorf("len(getAllPermutationsRec) = %d; want 6", len(perms))
	}

	initArray = []int{0, 1, 2, 3}
	perms = getAllPermutationsRec(initArray)
	if len(perms) != 24 {
		t.Errorf("len(getAllPermutationsRec) = %d; want 24", len(perms))
	}

	initArray = []int{0, 1, 2, 3, 4}
	perms = getAllPermutationsRec(initArray)
	if len(perms) != 120 {
		t.Errorf("len(getAllPermutationsRec) = %d; want 120", len(perms))
	}
}

func TestCalcAmpSignal(t *testing.T) {
	p := intcode.Program{}
	p.InitMemory("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0")
	phaseArr := []int{4, 3, 2, 1, 0}
	got := calcAmpSignal(&p, phaseArr)
	if got != 43210 {
		t.Errorf("calcAmpSignal = %d; want 43210", got)
	}

	p.Reset()
	p.InitMemory("3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0")
	phaseArr = []int{0, 1, 2, 3, 4}
	got = calcAmpSignal(&p, phaseArr)
	if got != 54321 {
		t.Errorf("calcAmpSignal = %d; want 54321", got)
	}

	p.Reset()
	p.InitMemory("3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0")
	phaseArr = []int{1, 0, 4, 3, 2}
	got = calcAmpSignal(&p, phaseArr)
	if got != 65210 {
		t.Errorf("calcAmpSignal = %d; want 65210", got)
	}
}
