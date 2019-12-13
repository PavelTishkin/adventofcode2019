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
	transferSteps := calcTransferSteps(skyMap["YOU"], skyMap["SAN"], skyMap["COM"])
	fmt.Printf("Part 2: %d\n", transferSteps)
}

func calcTransferSteps(obj1 *celBody, obj2 *celBody, root *celBody) int {
	path1 := getPathToRoot(obj1, root)
	path2 := getPathToRoot(obj2, root)
	lastCommon := getLastCommon(path1, path2)
	count1 := countStepsToNode(lastCommon, path1)
	count2 := countStepsToNode(lastCommon, path2)
	return count1 + count2
}

func countStepsToNode(node *celBody, path []*celBody) int {
	stepCount := 0
	for i := len(path) - 1; i > 0; i-- {
		if node.label == path[i].label {
			return stepCount
		}
		stepCount++
	}
	return stepCount
}

func getLastCommon(path1 []*celBody, path2 []*celBody) *celBody {
	for i, obj := range path1 {
		if obj.label != path2[i].label {
			return obj.parent
		}
	}
	return nil
}

func getPathToRoot(currObj *celBody, rootObj *celBody) []*celBody {
	var objPath []*celBody
	for ok := true; ok; ok = (currObj.label != rootObj.label) {
		objPath = append([]*celBody{currObj.parent}, objPath...)
		currObj = currObj.parent
	}

	return objPath
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
