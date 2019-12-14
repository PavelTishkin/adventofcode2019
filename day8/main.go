package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput(os.Args[1])
	width, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	height, err := strconv.ParseInt(os.Args[3], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	checksum := getLayerChecksum(int(width), int(height), input)
	fmt.Printf("Part 1: %d\n", checksum)
}

func getLayerChecksum(width int, height int, input string) int {
	inputArr := getIntArr(input)
	layers := getLayers(width, height, inputArr)
	countZero := -1
	checksum := 0
	for _, layer := range layers {
		currCountZero := countLayerDigits(layer, 0)
		if currCountZero < countZero || countZero == -1 {
			countZero = currCountZero
			checksum = countLayerDigits(layer, 1) * countLayerDigits(layer, 2)
		}
	}

	return checksum
}

func countLayerDigits(layer [][]int, digit int) int {
	count := 0
	for _, row := range layer {
		for _, col := range row {
			if col == digit {
				count++
			}
		}
	}

	return count
}

func getLayers(width int, height int, input []int) [][][]int {
	var layers [][][]int

	layerSize := width * height

	for layerNum := 0; layerNum < len(input); layerNum += layerSize {
		var layer [][]int
		for widthNum := 0; widthNum < layerSize; widthNum += width {
			row := input[layerNum+widthNum : layerNum+widthNum+width]
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}

	return layers
}

func getIntArr(input string) []int {
	var ret []int
	for _, char := range input {
		ret = append(ret, int(char-'0'))
	}
	return ret
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
