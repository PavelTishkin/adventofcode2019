package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput(os.Args[1])
	expected, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1: %d\n", calcAnswer(input, 12, 2))
	fmt.Printf("Part 2: %d\n", bruteforceAnswer(input, int(expected)))
}

func bruteforceAnswer(input string, expected int) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			answer := calcAnswer(input, i, j)
			if answer == expected {
				return i*100 + j
			}
		}
	}
	return 0
}

func calcAnswer(input string, noun int, verb int) int {
	memory := strToIntArray(input)
	memory[1] = noun
	memory[2] = verb
	processInput(memory)

	return memory[0]
}

func processInputStr(input string) string {
	memory := strToIntArray(input)
	processInput(memory)
	return intArrToString(memory)
}

func processInput(memory []int) {
	var keepGoing bool = true
	var nextIndex int = 0
	var nextOp, noun, verb, storePos int

	for keepGoing {
		nextOp = memory[nextIndex]

		switch nextOp {
		case 99:
			keepGoing = false
		case 1:
			noun = memory[nextIndex+1]
			verb = memory[nextIndex+2]
			storePos = memory[nextIndex+3]

			memory[storePos] = memory[noun] + memory[verb]
		case 2:
			noun = memory[nextIndex+1]
			verb = memory[nextIndex+2]
			storePos = memory[nextIndex+3]

			memory[storePos] = memory[noun] * memory[verb]
		}

		nextIndex += 4
	}
}

func strToIntArray(input string) []int {
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

	return intArray
}

func intArrToString(intArray []int) string {
	var strArray []string

	for _, intItem := range intArray {
		strItem := strconv.Itoa(intItem)

		strArray = append(strArray, strItem)
	}

	return strings.Join(strArray, ",")
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
