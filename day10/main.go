package main

import (
	"fmt"
	"math"
	"os"

	"../utils"
)

type point struct {
	x int
	y int
}

type vector struct {
	p        point
	angle    float64
	distance float64
}

func main() {
	lines := utils.ReadLines(os.Args[1])
	pointsArr := inputToPoints(lines)

	_, mostVisible := getBestAsteroid(pointsArr)
	fmt.Printf("Part 1: %d\n", mostVisible)
}

/*
getBestAsteroid returns point that has direct line of sight to most other asteroids
*/
func getBestAsteroid(pointArr []point) (point, int) {
	var bestPoint point
	bestVisible := 0

	for _, p := range pointArr {
		currVisible := calcVisibleAsteroids(p, pointArr)
		if currVisible > bestVisible {
			bestVisible = currVisible
			bestPoint = p
		}
	}

	return bestPoint, bestVisible
}

/*
calcVisibleAsteroids calculates number of asteroids directly visible
*/
func calcVisibleAsteroids(origin point, pointArr []point) int {
	visible := 0
	var angleMap = make(map[float64]vector)
	vectorsMap := createVectorsMap(origin, pointArr)
	for _, vector := range vectorsMap {
		if _, ok := angleMap[vector.angle]; !ok {
			angleMap[vector.angle] = vector
			visible++
		}
	}
	return visible
}

/*
createVectorsMap will iterate through array of points provided and
generate a list of vectors for each of the points relative to the origin point.
Each vector will contain angle and distance
*/
func createVectorsMap(origin point, pointArr []point) []vector {
	var vectorsArr []vector
	for _, p := range pointArr {
		if p != origin {
			relPoint := relPoint(origin, p)
			angle := calcVectorAngle(relPoint)
			// round angle to ten digits precision
			angle = math.Round(angle*float64(10)) / float64(10)
			vector := vector{
				p:        relPoint,
				angle:    angle,
				distance: calcVectorDistance(relPoint),
			}
			vectorsArr = append(vectorsArr, vector)
		}
	}
	return vectorsArr
}

/*
calculateVectorDistance will calculate distance between (0,0) and given point
*/
func calcVectorDistance(p point) float64 {
	return math.Sqrt(float64((p.x * p.x) + (p.y * p.y)))
}

/*
calcVectorAngle will calculate angle of a vector starting at (0,0) and ending at given point.
The angle is represented in a Cartesian coordinate system, with positive y axis being the start.
Angle is measured clockwise.
The asteroid map has an inverse y axis, but that shouldn't matter for the solution
*/
func calcVectorAngle(p point) float64 {
	angle := math.Atan2(float64(p.y), float64(p.x)) * 180.0 / math.Pi
	angle = math.Mod(90-angle+360, 360)
	return angle
}

/*
relPoint represents one point relative to another.
The first point is the origin, result is mapped as second point to first in relation to (0,0)
*/
func relPoint(p1 point, p2 point) point {
	retPoint := point{
		x: p2.x - p1.x,
		y: p2.y - p1.y,
	}

	return retPoint
}

/*
inputToPoints converts asterpod map into list of points.
Top left corner is (0,0), with each new character in a string increasing x coordinate.
Each new line is increasing y coordinate.
Asteroids are represented by '#', empty space represented by '.'
*/
func inputToPoints(textLines []string) []point {
	var pointArr []point
	for y, textLine := range textLines {
		for x, char := range textLine {
			if char == '#' {
				newPoint := point{
					x: x,
					y: y,
				}
				pointArr = append(pointArr, newPoint)
			}
		}
	}

	return pointArr
}
