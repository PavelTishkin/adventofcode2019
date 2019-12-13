package main

import (
	"testing"
)

func TestCalctransferSteps(t *testing.T) {
	skyMap := createMap([]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"})
	got := calcTransferSteps(skyMap["YOU"], skyMap["SAN"], skyMap["COM"])
	if got != 4 {
		t.Errorf("celBody.depCount = %v; want 4", 4)
	}
}

func TestCalcDependencies(t *testing.T) {
	skyMap := createMap([]string{"A)B", "B)C", "B)D"})
	calcDependencies(skyMap["A"])

	if skyMap["A"].depCount != 3 {
		t.Errorf("celBody.depCount = %v; want 3", skyMap["A"].depCount)
	}

	if skyMap["D"].depCount != 0 {
		t.Errorf("celBody.depCount = %v; want 0", skyMap["D"].depCount)
	}
}

func TestCreateMap(t *testing.T) {
	skyMap := createMap([]string{"A)B", "B)C", "B)D"})

	if skyMap["A"].parent != nil {
		t.Errorf("celBody.parent = %v; want nil", skyMap["A"].parent)
	}

	if len(skyMap["B"].children) != 2 {
		t.Errorf("len(celBody.children) = %d; want 2", len(skyMap["B"].children))
	}
}

func TestGetOrAddCelBody(t *testing.T) {
	var skyMap = make(map[string]*celBody)
	objA := &celBody{
		label: "A1",
	}
	skyMap["A"] = objA

	got := getOrAddCelBody(skyMap, "A")
	if got.label != "A1" {
		t.Errorf("celBody.label = %s; want A1", got.label)
	}

	got = getOrAddCelBody(skyMap, "B")
	if got.label != "B" {
		t.Errorf("celBody.label = %s; want B", got.label)
	}
}

func TestSplitOrbit(t *testing.T) {
	obj1, obj2 := splitOrbit("A)B")
	if obj1 != "A" && obj2 != "B" {
		t.Errorf("splitObjects = %s, %s; want A, B", obj1, obj2)
	}
}
