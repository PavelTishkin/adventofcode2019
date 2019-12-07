package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	minStr := os.Args[1]
	min, err := strconv.ParseInt(minStr, 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	maxStr := os.Args[2]
	max, err := strconv.ParseInt(maxStr, 10, 32)

	if err != nil {
		log.Fatal(err)
	}

	part1, part2 := countFilteredRange(int(min), int(max))
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func countFilteredRange(min int, max int) (int, int) {
	part1 := 0
	part2 := 0
	for i := min; i <= max; i++ {
		iStr := numToStrArray(i)
		if hasDouble(iStr) && isIncreasing(iStr) {
			part1++
			if hasTrueDouble(iStr) {
				part2++
			}
		}
	}
	return part1, part2
}

func hasTrueDouble(num []string) bool {
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			if (i == 0 || (i > 0 && num[i-1] != num[i])) && (i == len(num)-2 || (i < (len(num)-2) && num[i+2] != num[i])) {
				return true
			}
		}
	}
	return false
}

func hasDouble(num []string) bool {
	for i := 0; i < len(num)-1; i++ {
		if num[i] == num[i+1] {
			return true
		}
	}
	return false
}

func isIncreasing(num []string) bool {
	for i := 0; i < len(num)-1; i++ {
		if num[i] > num[i+1] {
			return false
		}
	}
	return true
}

func numToStrArray(num int) []string {
	strArr := []string{}
	bytes := []byte(strconv.Itoa(num))
	for _, b := range bytes {
		strArr = append(strArr, string(b))
	}
	return strArr
}
