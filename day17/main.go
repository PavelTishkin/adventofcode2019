package main

import (
	"fmt"
	"os"

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
	// fmt.Printf("Part 2: %d\n", timeToOxygen(droid))
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
