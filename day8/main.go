package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
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

	intArr := getIntArr(input)
	layers := getLayers(int(width), int(height), intArr)
	flatLayer := flattenLayers(int(width), int(height), layers)
	printLayer(flatLayer)
	convertLayerToPng(int(width), int(height), flatLayer, os.Args[4])
}

func printLayer(layer [][]int) {
	for _, row := range layer {
		for _, cell := range row {
			if cell == 0 {
				fmt.Print("*")
			} else if cell == 1 {
				fmt.Print("_")
			} else if cell == 2 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func convertLayerToPng(width int, height int, layer [][]int, imagePath string) {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	white := color.RGBA{0xff, 0xff, 0xff, 0xff}
	black := color.RGBA{0, 0, 0, 0xff}
	tranparent := color.RGBA{0, 0, 0, 0}

	for y, row := range layer {
		for x, cell := range row {
			switch cell {
			case 0:
				img.Set(x, y, black)
			case 1:
				img.Set(x, y, white)
			case 2:
				img.Set(x, y, tranparent)
			}

		}
	}

	// Encode as PNG.
	f, _ := os.Create(imagePath)
	png.Encode(f, img)
}

func flattenLayers(width int, height int, layers [][][]int) [][]int {
	var newLayer [][]int

	//Copy first layer as is
	firstLayer := layers[0]
	for _, row := range firstLayer {
		newRow := make([]int, len(row))
		copy(newRow, row)
		newLayer = append(newLayer, newRow)
	}

	for _, layer := range layers[1:] {
		for i, row := range layer {
			for j, cell := range row {
				if newLayer[i][j] == 2 && cell != 2 {
					newLayer[i][j] = cell
				}
			}
		}
	}

	return newLayer
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
