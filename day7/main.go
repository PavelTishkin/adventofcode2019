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
