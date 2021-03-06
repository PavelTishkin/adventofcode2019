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
	fmt.Printf("Part 1: %d\n", calcAnswer1(p))
	fmt.Printf("Part 2: %d\n", calcAnswer2(p))
}

func calcAnswer1(p intcode.Program) int {
	p.Reset()
	p.PushInput(1)
	p.Run()
	var outArr = p.GetOutput()
	return outArr[len(outArr)-1]
}

func calcAnswer2(p intcode.Program) int {
	p.Reset()
	p.PushInput(5)
	p.Run()
	var outArr = p.GetOutput()
	return outArr[len(outArr)-1]
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
