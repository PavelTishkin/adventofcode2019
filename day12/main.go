package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"../utils"
)

type point struct {
	x int
	y int
	z int
}

type moon struct {
	position *point
	velocity *point
}

func main() {
	input := utils.ReadLines(os.Args[1])
	moons := readMoons(input)
	timeForward(moons, 1000)
	systemEnergy := getSystemEnergy(moons)
	fmt.Printf("Part 1: %d\n", systemEnergy)
}

/*
getSystemEnergy calculates total energy of the system by adding all moons energies
*/
func getSystemEnergy(moons []*moon) int {
	systemEnergy := 0

	for _, m := range moons {
		systemEnergy += m.getTotalEnergy()
	}

	return systemEnergy
}

/*
getTotalEnergy calculates total energy of the moon as kinetic * potential
*/
func (m *moon) getTotalEnergy() int {
	return getEnergy(m.position) * getEnergy(m.velocity)
}

/*
getEnergy calculates energy of a point by adding absolute values
*/
func getEnergy(p *point) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)) + math.Abs(float64(p.z)))
}

/*
timeForward moves all the moons in the system n steps ahead
*/
func timeForward(moons []*moon, steps int) {
	for i := 0; i < steps; i++ {
		timeStep(moons)
	}
}

/*
timeStep iterates through each moon, calculates applied gravity and move moon to new position
*/
func timeStep(moons []*moon) {
	for _, m := range moons {
		m.applyGravity(moons)
	}

	for _, m := range moons {
		m.applyVelocity()
	}
}

/*
applyVelocity moves the moon to new position according to current velocity
*/
func (m *moon) applyVelocity() {
	m.position.x += m.velocity.x
	m.position.y += m.velocity.y
	m.position.z += m.velocity.z
}

/*
applyGravity updates velocity of the moon based on positions of other moons in the system
*/
func (m *moon) applyGravity(moons []*moon) {
	adjustedX, adjustedY, adjustedZ := 0, 0, 0

	for _, nMoon := range moons {
		adjustedX += calcGravityOffset(m.position.x, nMoon.position.x)
		adjustedY += calcGravityOffset(m.position.y, nMoon.position.y)
		adjustedZ += calcGravityOffset(m.position.z, nMoon.position.z)
	}

	m.velocity.x += adjustedX
	m.velocity.y += adjustedY
	m.velocity.z += adjustedZ
}

/*
calcGravityOffset compares coordinates of two points and return -1 if first coordinate is greater, 1 if smaller, 0 if equal
*/
func calcGravityOffset(c1, c2 int) int {
	if c1 < c2 {
		return 1
	} else if c1 > c2 {
		return -1
	}
	return 0
}

/*
readMoons converts string array of moon positions into array of moons with
initialized position and zero velocity vector
*/
func readMoons(lines []string) []*moon {
	var moons []*moon
	for _, line := range lines {
		re := regexp.MustCompile(`^<x=([^,]*), y=([^,]*), z=(.*)>$`)
		matched := re.FindAllStringSubmatch(line, -1)
		position := point{x: strToInt(matched[0][1]),
			y: strToInt(matched[0][2]),
			z: strToInt(matched[0][3]),
		}
		velocity := point{x: 0, y: 0, z: 0}
		newMoon := moon{
			position: &position,
			velocity: &velocity,
		}

		moons = append(moons, &newMoon)
	}

	return moons
}

func strToInt(strVal string) int {
	intVal, err := strconv.ParseInt(strVal, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return int(intVal)
}
