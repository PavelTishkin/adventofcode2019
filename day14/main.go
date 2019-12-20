package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"../utils"
)

type weightedElement struct {
	element string
	weight  int64
}

type reaction struct {
	input  []weightedElement
	output weightedElement
}

type beaker struct {
	reactions   []reaction
	reactionMap map[string]reaction
	leftovers   map[string]int64
	oreUsed     int64
}

func main() {
	input := utils.ReadLines(os.Args[1])
	oreUsed := calcOreForFuel(input)
	fmt.Printf("Part 1: %d\n", oreUsed)

	fuelProduced := calcFuelForOre(input)
	fmt.Printf("Part 2: %d\n", fuelProduced)
}

/*
calcFuelForOre calculates amount of FUEL produced with 1000000000000 ORE
*/
func calcFuelForOre(input []string) int64 {
	reactions := parseReactions(input)
	cleanBeaker := beaker{
		reactions:   reactions,
		reactionMap: createReactionMap(reactions),
		leftovers:   initLeftovers(reactions),
		oreUsed:     0,
	}

	oreLimit := int64(1000000000000)
	left := int64(1)
	right := oreLimit
	middle := int64(math.Floor(float64(left+right) / 2.0))

	for left <= right {
		cleanBeaker.reset()
		cleanBeaker.requestElement("FUEL", left)
		oreLeft := cleanBeaker.oreUsed
		cleanBeaker.reset()
		cleanBeaker.requestElement("FUEL", right)
		oreRight := cleanBeaker.oreUsed
		if left == right-1 && oreLeft < oreLimit && oreRight > oreLimit {
			return left
		}

		cleanBeaker.reset()
		cleanBeaker.requestElement("FUEL", middle)
		oreMiddle := cleanBeaker.oreUsed

		if oreMiddle < oreLimit {
			left = middle
		} else {
			right = middle
		}
		middle = int64(math.Floor(float64(left+right) / 2.0))
	}

	return 0
}

/*
calcOreForFuel calculates amount of ORE required to produce one unit of FUEL
*/
func calcOreForFuel(input []string) int64 {
	reactions := parseReactions(input)
	cleanBeaker := beaker{
		reactions:   reactions,
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
func (b *beaker) requestElement(elementRequired string, quantityRequired int64) {
	stillRequired := b.useElement(elementRequired, quantityRequired)
	if stillRequired > 0 {
		b.makeElement(elementRequired, stillRequired)
		b.useElement(elementRequired, stillRequired)
	}
}

/*
makeElement produces element (in multiples if required), based on reaction formula.
If ORE is used, it is counted up
*/
func (b *beaker) makeElement(elementRequired string, quantityRequired int64) {
	currReaction := b.reactionMap[elementRequired]
	quantityProduced := int64(math.Ceil(float64(quantityRequired) / float64(currReaction.output.weight)))
	for _, inputElement := range currReaction.input {
		if inputElement.element == "ORE" {
			b.oreUsed += inputElement.weight * quantityProduced
		} else {
			b.requestElement(inputElement.element, inputElement.weight*quantityProduced)
		}
	}
	b.leftovers[currReaction.output.element] += currReaction.output.weight * quantityProduced
}

/*
useElement reduces specific element leftover count by quantityRequired.
If there are not enough elements, return amount of elements still required
*/
func (b *beaker) useElement(elementRequired string, quantityRequired int64) int64 {
	if b.leftovers[elementRequired] >= quantityRequired {
		b.leftovers[elementRequired] -= quantityRequired
		return 0
	}

	stillRequired := quantityRequired - b.leftovers[elementRequired]
	b.leftovers[elementRequired] = 0
	return stillRequired
}

/*
initLeftovers initialises map of how many leftover elements there currently are
*/
func initLeftovers(reactions []reaction) map[string]int64 {
	leftovers := make(map[string]int64)

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

func (b *beaker) reset() {
	b.leftovers = initLeftovers(b.reactions)
	b.reactionMap = createReactionMap(b.reactions)
	b.oreUsed = 0
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
		weight:  int64(weight),
	}

	return el
}
