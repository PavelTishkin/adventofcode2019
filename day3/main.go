package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	p1 point
	p2 point
}

func main() {
	wire1, wire2 := readInput(os.Args[1])
	minDistance := getMinDistance(wire1, wire2)
	fmt.Printf("Part1: %d\n", minDistance)
	minWireDistance := getMinWireDistance(wire1, wire2)
	fmt.Printf("Part2: %d\n", minWireDistance)
}

func getMinDistance(wire1 string, wire2 string) int {
	wire1Points := strToPointArr(wire1)
	wire2Points := strToPointArr(wire2)

	minDistance := -1
	intersects := findIntersects(wire1Points, wire2Points)
	for _, intersect := range intersects {
		currDistance := getManhattanDistance(intersect)
		if minDistance == -1 || minDistance > currDistance {
			minDistance = currDistance
		}
	}

	return minDistance
}

func getMinWireDistance(wire1 string, wire2 string) int {
	wire1Points := strToPointArr(wire1)
	wire2Points := strToPointArr(wire2)

	minDistance := -1
	intersects := findIntersects(wire1Points, wire2Points)
	for _, intersect := range intersects {
		currDistance := getWireDistanceToIntersection(wire1Points, intersect) +
			getWireDistanceToIntersection(wire2Points, intersect)
		if minDistance == -1 || minDistance > currDistance {
			minDistance = currDistance
		}
	}

	return minDistance
}

func getWireDistanceToIntersection(wire []point, p point) int {
	fullWire := append([]point{point{x: 0, y: 0}}, wire...)
	currDist := 0
	for i := 0; i < len(fullWire)-1; i++ {
		l := line{p1: fullWire[i], p2: fullWire[i+1]}
		if isPointOnLine(l, p) {
			currDist += getPointsDistance(fullWire[i], p)
			return currDist
		}
		currDist += getPointsDistance(fullWire[i], fullWire[i+1])
	}
	return currDist
}

func isPointOnLine(l line, p point) bool {
	ol := orderLine(l)
	if isVertical(l) {
		if ol.p1.x == p.x && ol.p1.y <= p.y && ol.p2.y >= p.y {
			return true
		}
	}
	if ol.p1.y == p.y && ol.p1.x <= p.x && ol.p2.x >= p.x {
		return true
	}
	return false
}

func getPointsDistance(p1 point, p2 point) int {
	return getManhattanDistance(point{x: p1.x - p2.x, y: p1.y - p2.y})
}

func getManhattanDistance(p point) int {
	d := 0
	if p.x > 0 {
		d += p.x
	} else {
		d -= p.x
	}

	if p.y > 0 {
		d += p.y
	} else {
		d -= p.y
	}

	return d
}

func findIntersects(wire1Points []point, wire2Points []point) []point {
	intersects := []point{}
	for i, w1p := range wire1Points[:len(wire1Points)-1] {
		for j, w2p := range wire2Points[:len(wire2Points)-1] {
			intersect := getIntersect(line{p1: w1p, p2: wire1Points[i+1]}, line{p1: w2p, p2: wire2Points[j+1]})
			if intersect != nil {
				intersects = append(intersects, *intersect)
			}
		}
	}
	return intersects
}

func getIntersect(l1 line, l2 line) *point {
	if (isVertical(l1) && isVertical(l2)) || (isHorizontal(l1) && isHorizontal(l2)) {
		return nil
	}

	var ver line
	if isVertical(l1) {
		ver = l1
	} else {
		ver = l2
	}
	ver = orderLine(ver)

	var hor line
	if isHorizontal(l1) {
		hor = l1
	} else {
		hor = l2
	}
	hor = orderLine(hor)

	if hor.p1.x <= ver.p1.x && hor.p2.x >= ver.p1.x && ver.p1.y <= hor.p1.y && ver.p2.y >= hor.p1.y {
		return &point{x: ver.p1.x, y: hor.p1.y}
	}
	return nil
}

func orderLine(l line) line {
	if isHorizontal(l) {
		if l.p1.x < l.p2.x {
			return l
		}
		return line{p1: l.p2, p2: l.p1}
	}
	if l.p1.y < l.p2.y {
		return l
	}
	return line{p1: l.p2, p2: l.p1}
}

func isHorizontal(l line) bool {
	if l.p1.y == l.p2.y {
		return true
	}
	return false
}

func isVertical(l line) bool {
	if l.p1.x == l.p2.x {
		return true
	}
	return false
}

func strToPointArr(wire string) []point {
	pointArr := []point{point{x: 0, y: 0}}

	instructions := strings.Split(wire, ",")

	for _, instruction := range instructions {
		direction := string(instruction[0])
		d, err := strconv.ParseInt(string(instruction[1:]), 10, 32)
		distance := int(d)

		if err != nil {
			fmt.Println(instruction[1:])
			log.Fatal(err)
		}

		lastPoint := pointArr[len(pointArr)-1]

		switch direction {
		case "U":
			nextPoint := point{x: lastPoint.x, y: lastPoint.y + distance}
			pointArr = append(pointArr, nextPoint)
		case "D":
			nextPoint := point{x: lastPoint.x, y: lastPoint.y - distance}
			pointArr = append(pointArr, nextPoint)
		case "R":
			nextPoint := point{x: lastPoint.x + distance, y: lastPoint.y}
			pointArr = append(pointArr, nextPoint)
		case "L":
			nextPoint := point{x: lastPoint.x - distance, y: lastPoint.y}
			pointArr = append(pointArr, nextPoint)
		default:
			log.Fatal("Unknown case: " + direction)
		}
	}

	return pointArr[1:]
}

func readInput(filename string) (string, string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	wire1 := scanner.Text()
	scanner.Scan()
	wire2 := scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return wire1, wire2
}
