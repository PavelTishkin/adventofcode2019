package utils

import (
	"bufio"
	"log"
	"os"
)

/*
ReadLines will read file line by line and return an array of found strings
*/
func ReadLines(filename string) []string {
	var foundLines []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currLine := scanner.Text()

		foundLines = append(foundLines, currLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return foundLines
}
