package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"../utils"
)

func main() {
	input := utils.ReadLines(os.Args[1])[0]

	inputNum := stringToNumArray(input)
	pattern := []int64{0, 1, 0, -1}
	part1 := numArrayToString(fftFeedback(inputNum, pattern, 100)[:8])
	fmt.Printf("Part 1: %s\n", part1)
}

/*
fftFeedback calculates fft value in a feedback mode, repeatedly feeding output into input for the next function
*/
func fftFeedback(input, pattern []int64, times int) []int64 {
	fftInput := input
	expandedPattern := expandPattern(pattern, len(fftInput))
	for i := 0; i < times; i++ {
		fftInput = fft(fftInput, expandedPattern)
	}

	return fftInput
}

/*
fft calculates FFT value for input over pattern
*/
func fft(input []int64, expandedPattern [][]int64) []int64 {
	result := make([]int64, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = mulInputByPattern(input, expandedPattern[i])
	}

	return result
}

/*
mulInputByPattern itereates through input digits multiplying each digit by next number in pattern array.
Pattern array wraps around. Returns last digit of the added values of the result
*/
func mulInputByPattern(input, pattern []int64) int64 {
	result := int64(0)
	patternLen := len(pattern)

	for i, num := range input {
		result += num * pattern[i%patternLen]
	}

	return int64(math.Abs(float64(result))) % 10
}

/*
expandPattern expands given pattern according to FFT pattern rules and creates a multidimentional array of pattern values
*/
func expandPattern(pattern []int64, times int) [][]int64 {
	var patterns [][]int64

	for i := 0; i < times; i++ {
		var newPattern []int64
		for len(newPattern) < times+1 {
			for j := 0; j < len(pattern); j++ {
				for k := 0; k < i+1; k++ {
					newPattern = append(newPattern, pattern[j])
				}
			}
		}
		newPattern = newPattern[1 : times+1]
		patterns = append(patterns, newPattern)
	}
	return patterns
}

/*
stringToNumArray takes a number and separates into array of individual digits.
Using int64 just in case part2 will require to use big numbers
*/
func stringToNumArray(input string) []int64 {
	var numArr []int64
	for _, digit := range input {
		numArr = append(numArr, int64(digit-'0'))
	}

	return numArr
}

/*
numArrayToString converts array of ints to a string
*/
func numArrayToString(input []int64) string {
	var retVal string
	for _, num := range input {
		retVal += strconv.Itoa(int(num))
	}
	return retVal
}
