package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"../utils"
)

type weightedElement struct {
	element string
	weight  int
}

type reaction struct {
	input  []weightedElement
	output weightedElement
}

type beaker struct {
	reactionMap map[string]reaction
	leftovers   map[string]int
	oreUsed     int
}

func main() {
	input := utils.ReadLines(os.Args[1])
	oreUsed := calcOreForFuel(input)
	fmt.Printf("Part 1: %d\n", oreUsed)
}

/*
calcOreForFuel calculates amount of ore required to produce one unit of fuel
*/
func calcOreForFuel(input []string) int {
	reactions := parseReactions(input)
	cleanBeaker := beaker{
		reactionMap: createReactionMap(reactions),
		leftovers:   initLeftovers(reactions),
		oreUsed:     0,
	}

	cleanBeaker.requestElement("FUEL", 1)

	return cleanBeaker.oreUsed
}

/*
requestElement is a recursive function that takes name and quantity of an element and returns amount of ore required to produce it
*/
func (b *beaker) requestElement(elementRequired string, quantityRequired int) {
	for i := quantityRequired; i > 0; i-- {
		if b.hasElement(elementRequired) {
			b.useElement(elementRequired)
		} else {
			b.makeElement(elementRequired)
			b.useElement(elementRequired)
		}
	}
}

/*
makeElement produces element (in multiples if required), based on reaction formula.
If ORE is used, it is counted up
*/
func (b *beaker) makeElement(elementRequired string) {
	currReaction := b.reactionMap[elementRequired]
	for _, inputElement := range currReaction.input {
		if inputElement.element == "ORE" {
			b.oreUsed += inputElement.weight
		} else {
			b.requestElement(inputElement.element, inputElement.weight)
		}
	}
	b.leftovers[currReaction.output.element] += currReaction.output.weight
}

/*
hasElement returns true if leftover has at least one element
*/
func (b *beaker) hasElement(elementRequired string) bool {
	return b.leftovers[elementRequired] > 0
}

/*
useElement reduces specific element leftover count by 1
*/
func (b *beaker) useElement(elementRequired string) {
	b.leftovers[elementRequired]--
}

/*
initLeftovers initialises map of how many leftover elements there currently are
*/
func initLeftovers(reactions []reaction) map[string]int {
	leftovers := make(map[string]int)

	leftovers["ORE"] = 0
	for _, reaction := range reactions {
		leftovers[reaction.output.element] = 0
	}

	return leftovers
}

/*
createReactionMap creates a map with a key of output reaction
*/
func createReactionMap(reactions []reaction) map[string]reaction {
	reactionMap := make(map[string]reaction)

	for _, reaction := range reactions {
		reactionMap[reaction.output.element] = reaction
	}

	return reactionMap
}

/*
parseReaction reads a list of reactions and converts to objects
*/
func parseReactions(input []string) []reaction {
	var reactions []reaction

	for _, line := range input {
		var inputElements []weightedElement
		equation := strings.Split(line, "=>")
		elements := strings.Split(equation[0], ",")
		for _, element := range elements {
			we := parseElement(element)
			inputElements = append(inputElements, we)
		}
		outputElement := parseElement(equation[1])

		newReaction := reaction{
			input:  inputElements,
			output: outputElement,
		}

		reactions = append(reactions, newReaction)
	}

	return reactions
}

/*
parseElement reads a string and returns a weighted element
*/
func parseElement(element string) weightedElement {
	elItems := strings.Split(strings.Trim(element, " "), " ")
	weight, err := strconv.Atoi(elItems[0])
	if err != nil {
		log.Fatal(err)
	}

	el := weightedElement{
		element: elItems[1],
		weight:  weight,
	}

	return el
}
