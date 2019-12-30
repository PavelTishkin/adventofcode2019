package main

import (
	"fmt"
	"os"
	"strings"

	"../intcode"
	"../utils"
)

type point struct {
	x int
	y int
}

type droid struct {
	p             *intcode.Program
	scaffoldMap   [][]rune
	intersections []*point
}

func main() {
	input := utils.ReadLines(os.Args[1])[0]
	code := intcode.Program{}
	code.InitMemory(input)

	droid := initDroid(&code)
	droid.drawScaffold()
	//droid.printScafold()
	droid.getIntersections()
	fmt.Printf("Part 1: %d\n", droid.getIntersectionsSum())
	fmt.Printf("Part 2: %d\n", droid.navigateScaffolding())
}

/*
initDroid initializes droid with an intcode program
*/
func initDroid(p *intcode.Program) *droid {
	d := droid{
		p: p,
	}
	return &d
}

/*
drawScaffold reads values provided by droid to construct a scaffold map
*/
func (d *droid) drawScaffold() {
	var scaffoldMap [][]rune
	d.p.Run()
	scaffold := d.p.GetOutput()
	var row []rune
	for _, cell := range scaffold {
		nextChar := rune(cell)
		if cell != 10 {
			row = append(row, nextChar)
		} else {
			scaffoldMap = append(scaffoldMap, row)
			row = []rune{}
		}
	}

	d.scaffoldMap = scaffoldMap
}

func (d *droid) navigateScaffolding() int64 {
	d.p.Reset()
	d.p.SetMemoryValue(0, 2)
	main := createASCIIArrFromString("A,B,A,A,B,C,B,C,C,B")
	a := createASCIIArrFromString("L12,R8,L6,R8,L6")
	b := createASCIIArrFromString("R8,L12,L12,R8")
	c := createASCIIArrFromString("L6,R6,L12")
	d.inputInstructions(main, a, b, c)
	d.p.Run()
	output := d.p.GetOutput()
	return output[len(output)-1]
}

/*
getIntersections iterates through scaffold and finds location of all positions that have adjacent non-empty cells on all sides
*/
func (d *droid) getIntersections() {
	var intersections []*point

	for y, row := range d.scaffoldMap {
		for x := range row {
			if y > 0 && y < len(d.scaffoldMap)-1 && x > 0 && x < len(row)-1 {
				if d.isNotEmpty(x, y) && d.isNotEmpty(x-1, y) && d.isNotEmpty(x+1, y) && d.isNotEmpty(x, y-1) && d.isNotEmpty(x, y+1) {
					intersections = append(intersections, &point{x: x, y: y})
				}
			}
		}
	}

	d.intersections = intersections
}

/*
getIntersectionsSum adds product of x and y for all found intersections
*/
func (d *droid) getIntersectionsSum() int {
	total := 0

	for _, intersection := range d.intersections {
		total += intersection.x * intersection.y
	}

	return total
}

/*
isNotEmpty returns true if cell at position x,y is not a '.'
*/
func (d *droid) isNotEmpty(x, y int) bool {
	return d.scaffoldMap[y][x] != '.'
}

/*
inputInstructions provides instructions for the droid to navigate scaffolding
*/
func (d *droid) inputInstructions(main, a, b, c []int64) {
	input := append(main, a...)
	input = append(input, b...)
	input = append(input, c...)
	input = append(input, createASCIIArrFromString("n")...)
	d.p.SetInput(input)
}

/*
createASCIIArrFromString creates array of ascii characters ord numbers to pass to droid
*/
func createASCIIArrFromString(input string) []int64 {
	var output []int64
	inputArr := strings.Split(input, ",")

	for _, letter := range inputArr {
		output = append(output, int64(letter[0]))
		if len(letter) > 1 {
			output = append(output, int64(','))
			for _, digit := range letter[1:] {
				output = append(output, int64(digit))
			}
		}

		output = append(output, int64(','))
	}
	output[len(output)-1] = 10
	return output
}

/*
printScafold displays graphical map of scaffold for the droid
*/
func (d *droid) printScafold() {
	for _, row := range d.scaffoldMap {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Println()
	}
}
