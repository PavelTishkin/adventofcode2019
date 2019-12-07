package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if os.Args[1] == "1" {
		lines := readFile(os.Args[2])
		modTotal := 0
		for _, module := range lines {
			modTotal += calcMass(module)
		}
		fmt.Printf("Total weight: %d\n", modTotal)
	} else if os.Args[1] == "2" {
		lines := readFile(os.Args[2])
		modTotal := 0
		for _, module := range lines {
			modTotal += calcMassRec(module)
		}
		fmt.Printf("Total weight recursive: %d\n", modTotal)
	}

}

func calcMass(mass int) int {
	return int(mass/3) - 2
}

func calcMassRec(mass int) int {
	fuelReq := calcMass(mass)
	if fuelReq > 0 {
		fuelReq += calcMassRec(fuelReq)
	} else {
		return 0
	}

	return fuelReq
}

func readFile(filename string) []int {
	var lineArray []int
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine := scanner.Text()
		num, err := strconv.ParseInt(currLine, 10, 32)

		if err != nil {
			fmt.Println(currLine)
			log.Fatal(err)
		}

		lineArray = append(lineArray, int(num))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lineArray
}
