package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"../intcode"
	"../utils"
)

type point struct {
	x int
	y int
}

type droid struct {
	code           *intcode.Program // logic for the droid
	location       *point           // current location
	rootPath       *movementMap     // root of the movement map
	flatPath       []*movementMap   // all movements maps in an array
	visited        []*point         // list of points visited
	foundTarget    bool             // true if target is found
	oxygenPosition *point           // location of oxygen
}

type movementMap struct {
	location *point
	distance int

	parent     *movementMap
	childMoves []*movementMap
}

func main() {
	input := utils.ReadLines(os.Args[1])[0]
	p := intcode.Program{}
	p.InitMemory(input)

	droid := initDroid(&p)
	fmt.Printf("Part 1: %d\n", stepsToOxygen(droid))
	fmt.Printf("Part 2: %d\n", timeToOxygen(droid))
}

/*
initDroid initialises a new droid
*/
func initDroid(p *intcode.Program) *droid {
	// Initialize droid
	movementMapRoot := movementMap{location: &point{x: 0, y: 0}, distance: 0}
	d := droid{
		code:        p,
		location:    &point{x: 0, y: 0},
		rootPath:    &movementMapRoot,
		flatPath:    []*movementMap{&movementMapRoot},
		foundTarget: false,
	}
	d.code.Run()

	return &d
}

/*
stepsToOxygen calculates least amount of steps required to take to find a target in a blind maze
*/
func stepsToOxygen(d *droid) int {

	// Run search algorithm (assume no more than n steps are required)
	for i := 0; i < 1000; i++ {
		// Try to take next step from all existing moves with distance i
		tryMoves := d.getAllMovesWithDistance(i)
		for _, tryMove := range tryMoves {
			// First, navigate to the location of the move chosen
			d.moveToPoint(tryMove.location)
			// Take a step in every direction (will populate map of movements)
			d.moveEveryDirection()
		}
		// Check if target found flag is set, if yes, we're done
		if d.foundTarget {
			return i + 1
		}
	}

	log.Fatal("Could not find target...")
	return 0
}

/*
timeToOxygen calculates how much times it takes to fill up maze completely
*/
func timeToOxygen(d *droid) int {
	d.moveToPoint(d.oxygenPosition)

	// Reset initial location to oxygen point
	movementMapRoot := movementMap{location: &point{x: d.location.x, y: d.location.y}, distance: 0}
	d.rootPath = &movementMapRoot
	d.flatPath = []*movementMap{&movementMapRoot}

	timeToFill := 0
	tryMoves := d.getAllMovesWithDistance(timeToFill)

	for len(tryMoves) != 0 {
		for _, tryMove := range tryMoves {
			// First, navigate to the location of the move chosen
			d.moveToPoint(tryMove.location)
			// Take a step in every direction (will populate map of movements)
			d.moveEveryDirection()
		}
		timeToFill++
		tryMoves = d.getAllMovesWithDistance(timeToFill)
	}

	return timeToFill - 2
}

/*
getAllMovesWithDistance returns list of all movement maps that have specified distance
*/
func (d *droid) getAllMovesWithDistance(distance int) []*movementMap {
	var retMap []*movementMap
	for _, move := range d.flatPath {
		if move.distance == distance {
			retMap = append(retMap, move)
		}
	}

	return retMap
}

/*
getMovementMapByPoint returns a reference to a movementMap given a locatoin
*/
func (d *droid) getMovementMapByPoint(p *point) *movementMap {
	for _, move := range d.flatPath {
		if pointsEquals(move.location, p) {
			return move
		}
	}
	return nil
}

/*
moveToPoint moves droid back to a specified point by stepping back through map of previous steps
*/
func (d *droid) moveToPoint(p *point) {
	if !(d.location.x == p.x && d.location.y == p.y) {
		path := d.calculatePath(d.location, p)
		for _, step := range path {
			moveResult := d.move(step)
			if moveResult == 0 {
				log.Fatal("Move to point can't find path")
			}
		}
	}
}

/*
calculatePath calculates moves required to get from one point to the other, given a common root location
*/
func (d *droid) calculatePath(from, to *point) []int {
	path1 := d.backtraceToPoint(from, d.rootPath.location)
	path2 := d.backtraceToPoint(to, d.rootPath.location)
	lastCommon := getLastCommon(path1, path2)
	path1 = d.backtraceToPoint(from, lastCommon)
	path2 = d.backtraceToPoint(to, lastCommon)
	path2 = reversePointsArr(path2)
	combinedPath := append(path1, path2[1:]...)

	var moves []int
	for i, step := range combinedPath[1:] {
		moves = append(moves, getMoveDirection(combinedPath[i], step))
	}
	return moves
}

/*
backtraceToPoint finds sequence of points needed to step back to the point already visited
*/
func (d *droid) backtraceToPoint(currLoc, rootLoc *point) []*point {
	movement := d.findPointInPath(currLoc)
	backPath := []*point{movement.location}

	isEqual := pointsEquals(movement.location, rootLoc)

	for !isEqual {
		backPath = append(backPath, movement.parent.location)
		movement = movement.parent
		isEqual = pointsEquals(movement.location, rootLoc)
	}

	return backPath
}

/*
findPointInPath finds a point in movement map that led to the specified point
*/
func (d *droid) findPointInPath(p *point) *movementMap {
	for _, movement := range d.flatPath {
		if pointsEquals(movement.location, p) {
			return movement
		}
	}
	return nil
}

/*
getLastCommon returns closest common point in a tree of points between two paths
*/
func getLastCommon(path1, path2 []*point) *point {
	maxRange := int(math.Min(float64(len(path1)), float64(len(path2))))
	path1Mod := reversePointsArr(path1)
	path2Mod := reversePointsArr(path2)
	for i := 0; i < maxRange; i++ {
		if !pointsEquals(path1Mod[i], path2Mod[i]) {
			return path1Mod[i-1]
		}
	}
	return path1Mod[maxRange-1]
}

/*
reversePoints returns array of points in reverse order
*/
func reversePointsArr(points []*point) []*point {
	pointsRev := make([]*point, len(points))
	for i, point := range points {
		pointsRev[len(points)-1-i] = point
	}
	return pointsRev
}

/*
getMoveDirection returns int value of a move between two adjacent points.
Throws error if points are not adjacent
*/
func getMoveDirection(from, to *point) int {
	if from.x == to.x && from.y == to.y+1 {
		return 1
	} else if from.x == to.x && from.y == to.y-1 {
		return 2
	} else if from.x == to.x+1 && from.y == to.y {
		return 3
	} else if from.x == to.x-1 && from.y == to.y {
		return 4
	} else {
		log.Fatalf("Points %v and %v are not adjacent\n", from, to)
	}
	return 0
}

/*
moveEveryDirection attempts to move in all four directions and record results
*/
func (d *droid) moveEveryDirection() {
	currStep := d.getMovementMapByPoint(d.location)
	moveResult := d.move(1)
	if moveResult != 0 {
		if !d.hasVisitedCurrent() {
			d.addStep(currStep)
		}
		d.moveToPoint(currStep.location)
	}
	moveResult = d.move(2)
	if moveResult != 0 {
		if !d.hasVisitedCurrent() {
			d.addStep(currStep)
		}
		d.moveToPoint(currStep.location)
	}
	moveResult = d.move(3)
	if moveResult != 0 {
		if !d.hasVisitedCurrent() {
			d.addStep(currStep)
		}
		d.moveToPoint(currStep.location)
	}
	moveResult = d.move(4)
	if moveResult != 0 {
		if !d.hasVisitedCurrent() {
			d.addStep(currStep)
		}
		d.moveToPoint(currStep.location)
	}
}

/*
move command moved droid in a specified direction.
Updates current location.
direction values are 1 - north, 2 - south, 3 - west, 4 - east.
return values are 0 - hit a wall, 1 - moved successfully, 2 - found target
*/
func (d *droid) move(direction int) int {
	d.code.PushInput(int64(direction))
	d.code.Continue()

	moveResult := int(d.code.PopOutput())
	if moveResult != 0 {
		if direction == 1 {
			d.location.y--
		} else if direction == 2 {
			d.location.y++
		} else if direction == 3 {
			d.location.x--
		} else {
			d.location.x++
		}
	}

	if moveResult == 2 {
		d.foundTarget = true
		d.oxygenPosition = &point{x: d.location.x, y: d.location.y}
	}

	return moveResult
}

/*
addStep adds current droid location to a movement map
*/
func (d *droid) addStep(currStep *movementMap) {
	newPoint := point{x: d.location.x, y: d.location.y}
	newStep := movementMap{
		location: &newPoint,
		distance: currStep.distance + 1,
		parent:   currStep,
	}
	currStep.childMoves = append(currStep.childMoves, &newStep)
	d.flatPath = append(d.flatPath, &newStep)
}

/*
hasVisited returns true if droid has already visited its current location in the past
*/
func (d *droid) hasVisitedCurrent() bool {
	for _, visitedPoint := range d.visited {
		if d.location.x == visitedPoint.x && d.location.y == visitedPoint.y {
			return true
		}
	}
	return false
}

/*
pointsEquals returns true if x and y coordinates of two points are same
*/
func pointsEquals(p1, p2 *point) bool {
	return p1.x == p2.x && p1.y == p2.y
}
