package main

import (
	"strings"
	"testing"
)

func TestGetBestAsteroid(t *testing.T) {
	text := `.#..#
.....
#####
....#
...##`
	lines := strings.Split(text, "\n")
	pointsArr := inputToPoints(lines)

	bestAsteroid, mostVisible := getBestAsteroid(pointsArr)
	expectedAsteroid := point{3, 4}
	if bestAsteroid != expectedAsteroid {
		t.Errorf("getBestAsteroid.point = %v; expected %v", bestAsteroid, expectedAsteroid)
	}
	if mostVisible != 8 {
		t.Errorf("getBestAsteroid.mostVisible = %d; expected 8", mostVisible)
	}

	text = `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`
	lines = strings.Split(text, "\n")
	pointsArr = inputToPoints(lines)

	bestAsteroid, mostVisible = getBestAsteroid(pointsArr)
	expectedAsteroid = point{5, 8}
	if bestAsteroid != expectedAsteroid {
		t.Errorf("getBestAsteroid.p = %v; expected %v", bestAsteroid, expectedAsteroid)
	}
	if mostVisible != 33 {
		t.Errorf("getBestAsteroid.mostVisible = %d; expected 33", mostVisible)
	}

	text = `#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.`
	lines = strings.Split(text, "\n")
	pointsArr = inputToPoints(lines)

	bestAsteroid, mostVisible = getBestAsteroid(pointsArr)
	expectedAsteroid = point{1, 2}
	if bestAsteroid != expectedAsteroid {
		t.Errorf("getBestAsteroid.p = %v; expected %v", bestAsteroid, expectedAsteroid)
	}
	if mostVisible != 35 {
		t.Errorf("getBestAsteroid.mostVisible = %d; expected 35", mostVisible)
	}

	text = `.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..`
	lines = strings.Split(text, "\n")
	pointsArr = inputToPoints(lines)

	bestAsteroid, mostVisible = getBestAsteroid(pointsArr)
	expectedAsteroid = point{6, 3}
	if bestAsteroid != expectedAsteroid {
		t.Errorf("getBestAsteroid.p = %v; expected %v", bestAsteroid, expectedAsteroid)
	}
	if mostVisible != 41 {
		t.Errorf("getBestAsteroid.mostVisible = %d; expected 41", mostVisible)
	}

	text = `.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	lines = strings.Split(text, "\n")
	pointsArr = inputToPoints(lines)

	bestAsteroid, mostVisible = getBestAsteroid(pointsArr)
	expectedAsteroid = point{11, 13}
	if bestAsteroid != expectedAsteroid {
		t.Errorf("getBestAsteroid.p = %v; expected %v", bestAsteroid, expectedAsteroid)
	}
	if mostVisible != 210 {
		t.Errorf("getBestAsteroid.mostVisible = %d; expected 210", mostVisible)
	}
}

func TestCalcVisibleAsteroids(t *testing.T) {
	text := `.#..#
.....
#####
....#
...##`
	lines := strings.Split(text, "\n")
	pointsArr := inputToPoints(lines)

	originPoint := point{x: 3, y: 4}
	visibleAsteroids := calcVisibleAsteroids(originPoint, pointsArr)

	if visibleAsteroids != 8 {
		t.Errorf("calcVisibleAsteroids() = %d; expected 8", visibleAsteroids)
	}
}

func TestCreateVectorsArr(t *testing.T) {
	originPoint := point{x: 2, y: 1}
	pointArr := []point{
		point{x: 1, y: 2},
		point{x: 2, y: 1},
		point{x: -2, y: -1},
	}
	vectors := createVectorsMap(originPoint, pointArr)
	if len(vectors) != 2 {
		t.Errorf("len(vectors) = %d; expected 2", len(vectors))
	}
	expectedPoint := point{x: -1, y: 1}
	if vectors[0].p != expectedPoint {
		t.Errorf("vectors[0].p = %v; expected %v", vectors[0].p, expectedPoint)
	}
	if vectors[0].angle != 315 {
		t.Errorf("vectors[0].angle = %f; expected 315", vectors[0].angle)
	}
}

func TestCalcVectorDistance(t *testing.T) {
	got := calcVectorDistance(point{x: 3, y: 4})
	expected := float64(5)

	if got != expected {
		t.Errorf("got: %v; expected: %v", got, expected)
	}
}

func TestCalcVectorAngle(t *testing.T) {
	got := calcVectorAngle(point{x: 1, y: 1})
	expected := 45.0

	if got != expected {
		t.Errorf("Got: %f; expected: %f", got, expected)
	}

	got = calcVectorAngle(point{x: 1, y: 0})
	expected = 90.0

	if got != expected {
		t.Errorf("Got: %f; expected: %f", got, expected)
	}

	got = calcVectorAngle(point{x: 0, y: 1})
	expected = 0

	if got != expected {
		t.Errorf("Got: %f; expected: %f", got, expected)
	}

	got = calcVectorAngle(point{x: -1, y: -1})
	expected = 225

	if got != expected {
		t.Errorf("Got: %f; expected: %f", got, expected)
	}

	got = calcVectorAngle(point{x: -1, y: 1})
	expected = 315

	if got != expected {
		t.Errorf("Got: %f; expected: %f", got, expected)
	}
}

func TestInputToPoints(t *testing.T) {
	var lines []string
	lines = append(lines, ".#.")
	lines = append(lines, "#.#")
	lines = append(lines, "..#")

	got := inputToPoints(lines)

	var expected []point
	expected = append(expected, point{1, 0})
	expected = append(expected, point{0, 1})
	expected = append(expected, point{2, 1})
	expected = append(expected, point{2, 2})

	if !pointSlicesEqual(got, expected) {
		t.Errorf("got=%v; expected=%v", got, expected)
	}

}

func pointSlicesEqual(a, b []point) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v.x != b[i].x || v.y != b[i].y {
			return false
		}
	}
	return true
}
