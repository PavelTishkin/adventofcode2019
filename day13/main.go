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

func main() {
	input := utils.ReadLines(os.Args[1])
	p := intcode.Program{}
	p.InitMemory(input[0])

	field := mapField(&p)
	blocks := getObjByType(&field, 2)

	fmt.Printf("Part 1: %d\n", len(blocks))

	fmt.Printf("Part 2: %d\n", runArcadeToFinish(&p))
}

func runArcadeToFinish(p *intcode.Program) int {
	// Initialize arcade with input and run
	p.Reset()
	p.SetMemoryValue(0, int64(2))

	lastBallPos := point{x: 0, y: 0}
	currBallPos := point{x: 1, y: 1}

	for lastBallPos != currBallPos {
		// Read field
		field := mapField(p)

		// Get number of blocks
		numBlocks := len(getObjByType(&field, 2))
		if numBlocks == 0 {
			return getScore(&field)
		}

		// Get ball position
		ballPos := getObjByType(&field, 4)[0]
		lastBallPos = currBallPos
		currBallPos = ballPos

		// Get paddle position
		paddlePos := getObjByType(&field, 3)[0]
		// Move paddle with joystick
		if paddlePos.x < ballPos.x {
			p.PushInput(1)
		} else if paddlePos.x > ballPos.x {
			p.PushInput(-1)
		} else {
			p.PushInput(0)
		}
	}
	fmt.Println("Game over")

	return 0
}

/*
getObjByType returns locations of all objects with a given type on a map
*/
func getObjByType(field *map[point]int, objType int) []point {
	var locations []point

	for k, v := range *field {
		if v == objType {
			locations = append(locations, k)
		}
	}

	return locations
}

/*
getScore returns current score
*/
func getScore(field *map[point]int) int {
	for k, v := range *field {
		if k.x == -1 && k.y == 0 {
			return v
		}
	}

	return 0
}

/*
drawField initializes array of points that have some element in them according to an Intcode program
*/
func mapField(p *intcode.Program) map[point]int {
	fieldMap := make(map[point]int)

	p.Continue()

	fieldArr := p.GetOutput()

	for i := 0; i < len(fieldArr); i += 3 {
		newPoint := point{
			x: int(fieldArr[i]),
			y: int(fieldArr[i+1]),
		}
		fieldMap[newPoint] = int(fieldArr[i+2])
	}

	return fieldMap
}
