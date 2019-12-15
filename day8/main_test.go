package main

import "testing"

func TestFlattenLayer(t *testing.T) {
	intArr := getIntArr("0222112222120000")
	layers := getLayers(2, 2, intArr)
	layer := flattenLayers(2, 2, layers)
	expectedLayer := []int{0, 1}
	if !slicesEqual(layer[0], expectedLayer) {
		t.Errorf("first row = %v; expected = %v", layer[0], expectedLayer)
	}
	expectedLayer = []int{1, 0}
	if !slicesEqual(layer[1], expectedLayer) {
		t.Errorf("second row = %v; expected = %v", layer[1], expectedLayer)
	}
}

func TestGetLayerChecksum(t *testing.T) {
	got := getLayerChecksum(3, 2, "123456789012")
	expected := 1
	if got != expected {
		t.Errorf("getLayerChecksum = %d; expected = %d", got, expected)
	}
}

func TestGetIntArr(t *testing.T) {
	got := getIntArr("123456789012")
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2}
	if !slicesEqual(got, expected) {
		t.Errorf("getIntArr = %v; expected = %v", got, expected)
	}
}

func slicesEqual(a, b []int) bool {
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
