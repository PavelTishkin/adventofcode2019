package main

import "testing"

func TestMoonRun(t *testing.T) {
	lines := []string{"<x=-8, y=-10, z=0>",
		"<x=5, y=5, z=10>",
		"<x=2, y=-7, z=3>",
		"<x=9, y=-8, z=-3>"}
	moons := readMoons(lines)
	timeForward(moons, 100)

	expectedPos := &point{x: 16, y: -13, z: 23}
	expectedVel := &point{x: 7, y: 1, z: 1}
	systemEnergy := getSystemEnergy(moons)
	if !pointsEqual(moons[3].position, expectedPos) {
		t.Errorf("moons[3].position = %v; expected %v", moons[3].position, expectedPos)
	}
	if !pointsEqual(moons[3].velocity, expectedVel) {
		t.Errorf("moons[3].velocity = %v; expected %v", moons[3].velocity, expectedVel)
	}
	if systemEnergy != 1940 {
		t.Errorf("systemEnergy = %d, expected 1940", systemEnergy)
	}
}

func TestTimeStep(t *testing.T) {
	lines := []string{"<x=-1, y=0, z=2>",
		"<x=2, y=-10, z=-7>",
		"<x=4, y=-8, z=8>",
		"<x=3, y=5, z=-1>"}
	moons := readMoons(lines)
	timeStep(moons)

	expectedPos := &point{x: 2, y: -1, z: 1}
	expectedVel := &point{x: 3, y: -1, z: -1}
	if !pointsEqual(moons[0].position, expectedPos) {
		t.Errorf("moons[0].position = %v; expected %v", moons[0].position, expectedPos)
	}
	if !pointsEqual(moons[0].velocity, expectedVel) {
		t.Errorf("moons[0].velocity = %v; expected %v", moons[0].velocity, expectedVel)
	}

	timeStep(moons)
	timeStep(moons)

	expectedPos = &point{x: 2, y: 1, z: -5}
	expectedVel = &point{x: 1, y: 5, z: -4}
	if !pointsEqual(moons[2].position, expectedPos) {
		t.Errorf("moons[2].position = %v; expected %v", moons[2].position, expectedPos)
	}
	if !pointsEqual(moons[2].velocity, expectedVel) {
		t.Errorf("moons[2].velocity = %v; expected %v", moons[2].velocity, expectedVel)
	}

	timeForward(moons, 4)
	expectedPos = &point{x: 1, y: -4, z: -4}
	expectedVel = &point{x: -2, y: -4, z: -4}
	if !pointsEqual(moons[1].position, expectedPos) {
		t.Errorf("moons[1].position = %v; expected %v", moons[1].position, expectedPos)
	}
	if !pointsEqual(moons[1].velocity, expectedVel) {
		t.Errorf("moons[1].velocity = %v; expected %v", moons[1].velocity, expectedVel)
	}

	timeForward(moons, 3)
	systemEnergy := getSystemEnergy(moons)
	if systemEnergy != 179 {
		t.Errorf("systemEnergy = %d, expected 179", systemEnergy)
	}
}

func TestReadMoons(t *testing.T) {
	lines := []string{"<x=-15, y=10, z=-11>"}
	moons := readMoons(lines)

	if len(moons) != 1 {
		t.Errorf("len(moons) = %d; expected 1", len(moons))
	}
	if moons[0].position.x != -15 {
		t.Errorf("moons[0].position.x = %d; expected -15", moons[0].position.x)
	}
	if moons[0].position.y != 10 {
		t.Errorf("moons[0].position.y = %d; expected 10", moons[0].position.y)
	}
	if moons[0].position.z != -11 {
		t.Errorf("moons[0].position.z = %d; expected -15", moons[0].position.z)
	}
}

func pointsEqual(p1, p2 *point) bool {
	return p1.x == p2.x && p1.y == p2.y && p1.z == p2.z
}
