package main

import "testing"

func TestIsPointOnLine(t *testing.T) {
	got := isPointOnLine(line{point{-3, 1}, point{-3, 8}}, point{-3, 4})
	if !got {
		t.Error("isPointOnLine = false; want true")
	}

	got = isPointOnLine(line{point{5, 2}, point{-8, 2}}, point{-1, 2})
	if !got {
		t.Error("isPointOnLine = false; want true")
	}

	got = isPointOnLine(line{point{-3, 1}, point{-3, 8}}, point{1, 4})
	if got {
		t.Error("isPointOnLine = true; want false")
	}

	got = isPointOnLine(line{point{5, 2}, point{-8, 2}}, point{-1, 1})
	if got {
		t.Error("isPointOnLine = true; want false")
	}

	got = isPointOnLine(line{point{5, 2}, point{-8, 2}}, point{5, 2})
	if !got {
		t.Error("isPointOnLine = false; want true")
	}

	got = isPointOnLine(line{point{5, 2}, point{-8, 2}}, point{6, 2})
	if got {
		t.Error("isPointOnLine = true; want false")
	}
}

func TestPointsDistance(t *testing.T) {
	got := getPointsDistance(point{x: 3, y: 4}, point{x: 7, y: 12})
	if got != 12 {
		t.Errorf("getPointsDistance = %d; want 12", got)
	}

	got = getPointsDistance(point{x: -2, y: -5}, point{x: 7, y: 12})
	if got != 26 {
		t.Errorf("getPointsDistance = %d; want 26", got)
	}

	got = getPointsDistance(point{x: -2, y: -5}, point{x: -8, y: -4})
	if got != 7 {
		t.Errorf("getPointsDistance = %d; want 7", got)
	}
}

func TestManhattanDistance(t *testing.T) {
	got := getManhattanDistance(point{17, 2})
	if got != 19 {
		t.Errorf("getIntersect = %d; want 19", got)
	}

	got = getManhattanDistance(point{-3, 14})
	if got != 17 {
		t.Errorf("getIntersect = %d; want 17", got)
	}

	got = getManhattanDistance(point{2, -13})
	if got != 15 {
		t.Errorf("getIntersect = %d; want 15", got)
	}

	got = getManhattanDistance(point{0, 0})
	if got != 0 {
		t.Errorf("getIntersect = %d; want 0", got)
	}
}

func TestGetIntersect(t *testing.T) {
	got := getIntersect(line{point{x: 3, y: 3}, point{x: 3, y: 8}},
		line{point{x: -1, y: 4}, point{x: 11, y: 4}})
	expected := point{x: 3, y: 4}
	if *got != expected {
		t.Errorf("getIntersect = %v; want {3, 4}", got)
	}

	got = getIntersect(line{point{x: 3, y: 3}, point{x: 3, y: 8}},
		line{point{x: -1, y: 12}, point{x: 11, y: 12}})
	if got != nil {
		t.Errorf("getIntersect = %v; want nil", got)
	}

	got = getIntersect(line{point{x: 3, y: 3}, point{x: 3, y: 8}},
		line{point{x: 3, y: 3}, point{x: 11, y: 3}})
	expected = point{x: 3, y: 3}
	if *got != expected {
		t.Errorf("getIntersect = %v; want {3, 3}", got)
	}

	got = getIntersect(line{point{x: 3, y: 3}, point{x: 3, y: 8}},
		line{point{x: -1, y: 8}, point{x: 3, y: 8}})
	expected = point{x: 3, y: 8}
	if *got != expected {
		t.Errorf("getIntersect = %v; want {3, 8}", got)
	}
}

func TestOrderLine(t *testing.T) {
	got := orderLine(line{point{x: 5, y: 7}, point{x: 5, y: 12}})
	expected := line{point{x: 5, y: 7}, point{x: 5, y: 12}}
	if got != expected {
		t.Errorf("orderLine = %v; want {{5, 7}, {5, 12}}", got)
	}

	got = orderLine(line{point{x: 5, y: 12}, point{x: 5, y: 7}})
	expected = line{point{x: 5, y: 7}, point{x: 5, y: 12}}
	if got != expected {
		t.Errorf("orderLine = %v; want {{5, 7}, {5, 12}}", got)
	}

	got = orderLine(line{point{x: -1, y: 6}, point{x: 4, y: 6}})
	expected = line{point{x: -1, y: 6}, point{x: 4, y: 6}}
	if got != expected {
		t.Errorf("orderLine = %v; want {{-1, 6}, {4, 6}}", got)
	}

	got = orderLine(line{point{x: 4, y: 6}, point{x: -1, y: 6}})
	expected = line{point{x: -1, y: 6}, point{x: 4, y: 6}}
	if got != expected {
		t.Errorf("orderLine = %v; want {{-1, 6}, {4, 6}}", got)
	}
}

func TestIsHorizontal(t *testing.T) {
	got := isHorizontal(line{point{x: 5, y: 7}, point{x: 5, y: 12}})
	expected := false
	if got != expected {
		t.Error("isVertical = true; want false")
	}

	got = isHorizontal(line{point{x: -4, y: 12}, point{x: 5, y: 12}})
	expected = true
	if got != expected {
		t.Error("isVertical = false; want true")
	}
}

func TestIsVertical(t *testing.T) {
	got := isVertical(line{point{x: 5, y: 7}, point{x: 5, y: 12}})
	expected := true
	if got != expected {
		t.Error("isVertical = false; want true")
	}

	got = isVertical(line{point{x: -4, y: 12}, point{x: 5, y: 12}})
	expected = false
	if got != expected {
		t.Error("isVertical = true; want false")
	}
}

func TestStrToPointArr(t *testing.T) {
	got := strToPointArr("U5")
	expected := point{x: 0, y: 5}
	if got[0] != expected {
		t.Errorf("strToPointArr = %v; want {0, 5}", got)
	}

	got = strToPointArr("D2")
	expected = point{x: 0, y: -2}
	if got[0] != expected {
		t.Errorf("strToPointArr = %v; want {0, -2}", got)
	}

	got = strToPointArr("L4")
	expected = point{x: -4, y: 0}
	if got[0] != expected {
		t.Errorf("strToPointArr = %v; want {-4, 0}", got)
	}

	got = strToPointArr("R12")
	expected = point{x: 12, y: 0}
	if got[0] != expected {
		t.Errorf("strToPointArr = %v; want {12, 0}", got)
	}

	got = strToPointArr("D7,R5,U13")
	expectedArr := []point{point{x: 0, y: -7}, point{x: 5, y: -7}, point{x: 5, y: 6}}
	if got[0] != expectedArr[0] || got[1] != expectedArr[1] || got[2] != expectedArr[2] {
		t.Errorf("strToPointArr = %v; want {{0, -7}, {5, -7}, {5, 6}}", got)
	}
}
