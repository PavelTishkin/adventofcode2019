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
	p.PushInput(1)
	p.Run()
	boostKeycode := p.PopOutput()
	fmt.Printf("Part 1: %d\n", boostKeycode)
	p.Reset()
	p.PushInput(2)
	p.Run()
	boostKeycode = p.PopOutput()
	fmt.Printf("Part 2: %d\n", boostKeycode)
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
