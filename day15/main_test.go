package main

import "testing"

func TestGetAllMovesWithDistance(t *testing.T) {
	d := initTestDroid()
	gotMoves := d.getAllMovesWithDistance(1)
	if len(gotMoves) != 3 {
		t.Errorf("d.getAllMovesWithDistance = %d, expected 3", len(gotMoves))
	}

	gotMoves = d.getAllMovesWithDistance(3)
	if len(gotMoves) != 1 {
		t.Errorf("d.getAllMovesWithDistance = %d, expected 1", len(gotMoves))
	}
	expectedPoint := point{x: 2, y: 1}
	if !pointsEquals(&expectedPoint, gotMoves[0].location) {
		t.Errorf("d.getAllMovesWithDistance[0].location = %v, expected %v", gotMoves[0].location, expectedPoint)
	}

	gotMoves = d.getAllMovesWithDistance(0)
	if len(gotMoves) != 1 {
		t.Errorf("d.getAllMovesWithDistance = %d, expected 1", len(gotMoves))
	}
	expectedPoint = point{x: 0, y: 0}
	if !pointsEquals(&expectedPoint, gotMoves[0].location) {
		t.Errorf("d.getAllMovesWithDistance[0].location = %v, expected %v", gotMoves[0].location, expectedPoint)
	}
}

func TestCalculatePath(t *testing.T) {
	d := initTestDroid()

	p1 := point{x: 2, y: 1}
	p2 := point{x: 1, y: -1}
	gotPath := d.calculatePath(&p1, &p2)
	expectedPath := []int{3, 1, 1}
	if !intSlicesEqual(gotPath, expectedPath) {
		t.Errorf("calculatePath = %v, expected %v\n", gotPath, expectedPath)
	}

	p1 = point{x: 2, y: 1}
	p2 = point{x: 1, y: 0}
	gotPath = d.calculatePath(&p1, &p2)
	expectedPath = []int{3, 1}
	if !intSlicesEqual(gotPath, expectedPath) {
		t.Errorf("calculatePath = %v, expected %v\n", gotPath, expectedPath)
	}

	p1 = point{x: 2, y: 1}
	p2 = point{x: 0, y: 0}
	gotPath = d.calculatePath(&p1, &p2)
	expectedPath = []int{3, 1, 3}
	if !intSlicesEqual(gotPath, expectedPath) {
		t.Errorf("calculatePath = %v, expected %v\n", gotPath, expectedPath)
	}

	p1 = point{x: -1, y: 0}
	p2 = point{x: 2, y: 1}
	gotPath = d.calculatePath(&p1, &p2)
	expectedPath = []int{4, 4, 2, 4}
	if !intSlicesEqual(gotPath, expectedPath) {
		t.Errorf("calculatePath = %v, expected %v\n", gotPath, expectedPath)
	}

	p1 = point{x: 0, y: -1}
	p2 = point{x: 0, y: -1}
	gotPath = d.calculatePath(&p1, &p2)
	expectedPath = []int{}
	if len(gotPath) != 0 {
		t.Errorf("len(calculatePath) = %d, expected 0\n", len(gotPath))
	}
	if !intSlicesEqual(gotPath, expectedPath) {
		t.Errorf("calculatePath = %v, expected %v\n", gotPath, expectedPath)
	}
}

func TestBacktraceToPoint(t *testing.T) {
	d := initTestDroid()
	backtracePath := d.backtraceToPoint(&point{x: 2, y: 1}, &point{x: 1, y: 0})
	if len(backtracePath) != 3 {
		t.Errorf("len(d.backtraceToPoint) = %d, expected 3", len(backtracePath))
	}
	expectedPoint := point{x: 2, y: 1}
	if !pointsEquals(backtracePath[0], &expectedPoint) {
		t.Errorf("d.backtraceToPoint[0] = %v, expected %v", backtracePath[0], expectedPoint)
	}
	expectedPoint = point{x: 1, y: 1}
	if !pointsEquals(backtracePath[1], &expectedPoint) {
		t.Errorf("d.backtraceToPoint[0] = %v, expected %v", backtracePath[1], expectedPoint)
	}

	backtracePath = d.backtraceToPoint(&point{x: 1, y: 1}, &point{x: 1, y: 1})
	if len(backtracePath) != 1 {
		t.Errorf("len(d.backtraceToPoint) = %d, expected 1", len(backtracePath))
	}

	backtracePath = d.backtraceToPoint(&point{x: 1, y: 0}, &point{x: 0, y: 0})
	if len(backtracePath) != 2 {
		t.Errorf("len(d.backtraceToPoint) = %d, expected 2", len(backtracePath))
	}
}

func TestGetLastCommon(t *testing.T) {
	var path1 []*point
	var path2 []*point
	p1 := point{x: 0, y: 0}
	p2 := point{x: 1, y: 0}
	p3 := point{x: 1, y: -1}
	p4 := point{x: 1, y: 1}
	p5 := point{x: 2, y: 1}

	path1 = []*point{&p3, &p2, &p1}
	path2 = []*point{&p5, &p4, &p2, &p1}

	gotLastCommon := getLastCommon(path1, path2)
	if !pointsEquals(&p2, gotLastCommon) {
		t.Errorf("getLastCommon = %v, expected %v", gotLastCommon, p2)
	}

	path1 = []*point{&p3, &p2}
	path2 = []*point{&p5, &p4, &p2}

	gotLastCommon = getLastCommon(path1, path2)
	if !pointsEquals(&p2, gotLastCommon) {
		t.Errorf("getLastCommon = %v, expected %v", gotLastCommon, p2)
	}

	path1 = []*point{&p2, &p1}
	path2 = []*point{&p5, &p4, &p2, &p1}

	gotLastCommon = getLastCommon(path1, path2)
	if gotLastCommon == nil {
		t.Errorf("getLastCommon = nil, expected %v", p2)
	}
	if !pointsEquals(&p2, gotLastCommon) {
		t.Errorf("getLastCommon = %v, expected %v", gotLastCommon, p2)
	}

	path1 = []*point{&p5, &p4, &p2, &p1}
	path2 = []*point{&p2, &p1}

	gotLastCommon = getLastCommon(path1, path2)
	if gotLastCommon == nil {
		t.Errorf("getLastCommon = nil, expected %v", p2)
	}
	if !pointsEquals(&p2, gotLastCommon) {
		t.Errorf("getLastCommon = %v, expected %v", gotLastCommon, p2)
	}

	path1 = []*point{&p5, &p4, &p2, &p1}
	path2 = []*point{&p1}

	gotLastCommon = getLastCommon(path1, path2)
	if gotLastCommon == nil {
		t.Errorf("getLastCommon = nil, expected %v", p1)
	}
	if !pointsEquals(&p1, gotLastCommon) {
		t.Errorf("getLastCommon = %v, expected %v", gotLastCommon, p1)
	}
}

func TestReversePointsArr(t *testing.T) {
	var testArray []*point
	p1 := point{x: 1, y: -1}
	p2 := point{x: 2, y: 0}
	p3 := point{x: -1, y: 1}
	testArray = append(testArray, &p1)
	testArray = append(testArray, &p2)
	testArray = append(testArray, &p3)

	revArray := reversePointsArr(testArray)
	if len(revArray) != 3 {
		t.Errorf("reversePointsArr = %d, expected 3", len(revArray))
	}
	if !pointsEquals(revArray[0], &p3) {
		t.Errorf("reversePointsArr[0] = %v, expected %v", revArray[0], p3)
	}
}

func TestGetMoveDirection(t *testing.T) {
	p1 := point{x: 0, y: 0}
	p2 := point{x: 0, y: -1}
	p3 := point{x: 0, y: 1}
	p4 := point{x: -1, y: 0}
	p5 := point{x: 1, y: 0}

	gotDirection := getMoveDirection(&p1, &p2)
	if gotDirection != 1 {
		t.Errorf("getMoveDirection(%v, %v) = %d, expected 1", p1, p2, gotDirection)
	}
	gotDirection = getMoveDirection(&p1, &p3)
	if gotDirection != 2 {
		t.Errorf("getMoveDirection(%v, %v) = %d, expected 2", p1, p3, gotDirection)
	}
	gotDirection = getMoveDirection(&p1, &p4)
	if gotDirection != 3 {
		t.Errorf("getMoveDirection(%v, %v) = %d, expected 3", p1, p4, gotDirection)
	}
	gotDirection = getMoveDirection(&p1, &p5)
	if gotDirection != 4 {
		t.Errorf("getMoveDirection(%v, %v) = %d, expected 4", p1, p5, gotDirection)
	}
}

func initTestDroid() *droid {
	movementMapRoot := movementMap{location: &point{x: 0, y: 0}, distance: 0}
	testDroid := droid{
		code:        nil,
		location:    &point{x: 0, y: 0},
		rootPath:    &movementMapRoot,
		flatPath:    []*movementMap{&movementMapRoot},
		foundTarget: false,
	}

	testDroid.location.x = -1
	testDroid.location.y = 0
	testDroid.addStep(&movementMapRoot)
	testDroid.location.x = 0
	testDroid.location.y = -1
	testDroid.addStep(&movementMapRoot)
	testDroid.location.x = 1
	testDroid.location.y = 0
	testDroid.addStep(&movementMapRoot)

	currStep := testDroid.getMovementMapByPoint(&point{x: 1, y: 0})
	testDroid.location.x = 1
	testDroid.location.y = -1
	testDroid.addStep(currStep)
	testDroid.location.x = 1
	testDroid.location.y = 1
	testDroid.addStep(currStep)

	currStep = testDroid.getMovementMapByPoint(&point{x: 1, y: 1})
	testDroid.location.x = 2
	testDroid.location.y = 1
	testDroid.addStep(currStep)

	return &testDroid
}

func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
