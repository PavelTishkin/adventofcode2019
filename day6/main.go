package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type celBody struct {
	label    string
	parent   *celBody
	children []*celBody
	depCount int
}

func main() {
	lines := readFile(os.Args[1])
	skyMap := createMap(lines)
	calcDependencies(skyMap["COM"])
	fmt.Printf("Part 1: %d\n", countAllDeps(skyMap))
}

func countAllDeps(skyMap map[string]*celBody) int {
	totalDepCount := 0
	for _, obj := range skyMap {
		totalDepCount += obj.depCount
	}

	return totalDepCount
}

func calcDependencies(rootObj *celBody) int {
	deps := 0
	for _, child := range rootObj.children {
		deps++
		deps += calcDependencies(child)
	}
	rootObj.depCount = deps
	return deps
}

func createMap(orbits []string) map[string]*celBody {
	var skyMap = make(map[string]*celBody)

	for _, orbit := range orbits {
		rootLabel, childLabel := splitOrbit(orbit)
		rootRef := getOrAddCelBody(skyMap, rootLabel)
		childRef := getOrAddCelBody(skyMap, childLabel)
		childRef.parent = rootRef
		rootRef.children = append(rootRef.children, childRef)
	}

	return skyMap
}

func getOrAddCelBody(skyMap map[string]*celBody, objLabel string) *celBody {
	if objRef, ok := skyMap[objLabel]; ok {
		return objRef
	}

	objRef := &celBody{
		label: objLabel,
	}
	skyMap[objLabel] = objRef
	return objRef
}

func splitOrbit(objDep string) (string, string) {
	strArr := strings.Split(objDep, ")")
	return strArr[0], strArr[1]
}

func readFile(filename string) []string {
	var objArray []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine := scanner.Text()

		objArray = append(objArray, currLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return objArray
}
