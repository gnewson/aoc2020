package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stretchr/stew/slice"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newline := scanner.Text()
		if newline == "" {
			lines = append(lines, line)
			line = ""
		} else {
			line += newline
		}
	}

	lines = append(lines, strings.TrimSpace(line))
	return lines, scanner.Err()
}

func countOccurences(lines []string) int {
	var count int
	for _, line := range lines {
		//fmt.Println(line)
		var letters []rune
		for _, letter := range line[:] {

			if !slice.Contains(letters, letter) {
				letters = append(letters, letter)
			}
		}
		count += len(letters)
		fmt.Println(count)
	}
	return count
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Number of answers: ", countOccurences(lines))
}
