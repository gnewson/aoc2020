package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, []int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lines []string
	var line string
	var sizes []int
	var size int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newline := scanner.Text()
		if newline == "" {
			lines = append(lines, line)
			sizes = append(sizes, size)
			line = ""
			size = 0
		} else {
			line += newline
			size++
		}
	}

	sizes = append(sizes, size)
	lines = append(lines, strings.TrimSpace(line))
	return lines, sizes, scanner.Err()
}

func countOccurences(lines []string, partySize []int) int {
	var count int
	for i, line := range lines {
		//fmt.Println(line)
		var letters map[rune]int = make(map[rune]int)
		for _, letter := range line[:] {
			letters[letter]++
		}

		size := partySize[i]
		for k, _ := range letters {
			if letters[k] == size {
				count++
			}
		}
	}
	return count
}

func main() {
	lines, parties, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("Number of answers: ", countOccurences(lines, parties))
}
