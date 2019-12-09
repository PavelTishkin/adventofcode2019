package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"../intcode"
)

func main() {
	input := readInput(os.Args[1])
	expected, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	var p = intcode.Program{}
	p.InitMemory(input)
	fmt.Printf("Part 1: %d\n", calcAnswer(p, 12, 2))
	fmt.Printf("Part 2: %d\n", bruteforceAnswer(p, int(expected)))
}

func bruteforceAnswer(p intcode.Program, expected int) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			answer := calcAnswer(p, i, j)
			if answer == expected {
				return i*100 + j
			}
		}
	}
	return 0
}

func calcAnswer(p intcode.Program, noun int, verb int) int {
	p.Reset()
	p.SetInput(noun, verb)
	p.Run()
	return p.GetOutput()
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
