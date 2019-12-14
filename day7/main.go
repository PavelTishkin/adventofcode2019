package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"../intcode"
)

func main() {
	input := readInput(os.Args[1])
	var p = intcode.Program{}
	p.InitMemory(input)
	initArray := []int{0, 1, 2, 3, 4}
	maxPhase := getMaxPhaseValue(&p, initArray)
	fmt.Printf("Part 1: %d\n", maxPhase)
	initFeedbackArray := []int{5, 6, 7, 8, 9}
	maxPhaseFeedback := getMaxPhaseFeedbackValue(&p, initFeedbackArray)
	fmt.Printf("Part 2: %d\n", maxPhaseFeedback)
}

func getMaxPhaseValue(p *intcode.Program, initArray []int) int {
	perms := getAllPermutationsRec(initArray)

	maxPhaseValue := 0
	for _, perm := range perms {
		currPhaseValue := calcAmpSignal(p, perm)
		if currPhaseValue > maxPhaseValue {
			maxPhaseValue = currPhaseValue
		}
	}

	return maxPhaseValue
}

func getMaxPhaseFeedbackValue(p *intcode.Program, initArray []int) int {
	perms := getAllPermutationsRec(initArray)

	maxPhaseValue := 0
	for _, perm := range perms {

		ampArray := initAmpArray(perm, p.GetMemory())
		currPhaseValue := calcAmpFeedbackSignal(ampArray)
		if currPhaseValue > maxPhaseValue {
			maxPhaseValue = currPhaseValue
		}
	}

	return maxPhaseValue
}

func initAmpArray(initArray []int, initMemory []int) []*intcode.Program {
	var ampArray []*intcode.Program

	for _, initVal := range initArray {
		var currAmp = intcode.Program{}
		currAmp.CopyMemory(initMemory)
		currAmp.PushInput(initVal)
		currAmp.Run()
		ampArray = append(ampArray, &currAmp)
	}

	return ampArray
}

func getAllPermutationsRec(initArray []int) [][]int {
	permArray := getAllPermutations(initArray, []int{}, [][]int{})
	return permArray
}

func getAllPermutations(initArray []int, buildArray []int, recursiveArray [][]int) [][]int {
	var retArr [][]int

	if len(initArray) == 1 {
		cpyBuildArray := make([]int, len(buildArray))
		copy(cpyBuildArray, buildArray)
		cpyBuildArray = append(cpyBuildArray, initArray[0])
		retArr = append(recursiveArray, cpyBuildArray)
	} else {
		for i := 0; i < len(initArray); i++ {
			cpyInitArray := make([]int, len(initArray))
			cpyBuildArray := make([]int, len(buildArray))
			copy(cpyInitArray, initArray)
			copy(cpyBuildArray, buildArray)
			cpyBuildArray = append(cpyBuildArray, cpyInitArray[i])
			newInitArray := append(cpyInitArray[:i], cpyInitArray[i+1:]...)
			retArr = append(retArr, getAllPermutations(newInitArray, cpyBuildArray, recursiveArray)...)
		}
	}

	return retArr
}

func calcAmpSignal(p *intcode.Program, ampArray []int) int {
	phaseVal := 0
	for _, val := range ampArray {
		p.Reset()
		p.PushInput(val)
		p.PushInput(phaseVal)
		p.Run()
		phaseVal = p.PopOutput()
	}
	return phaseVal
}

func calcAmpFeedbackSignal(ampArray []*intcode.Program) int {
	phaseVal := 0
	lastAmp := ampArray[len(ampArray)-1]
	for lastAmp.IsRunning() {
		for _, amp := range ampArray {
			//p.Reset()
			amp.PushInput(phaseVal)
			amp.Continue()
			phaseVal = amp.PopOutput()
		}
	}
	return phaseVal
}

func readInput(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	foundLine := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return foundLine
}
