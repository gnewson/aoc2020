package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type slope struct {
	across int
	down   int
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func calculateTrees(rows []string, step slope) int {
	width := len(rows[0])
	pos := 0
	var count int

	for i, row := range rows {
		if i%step.down == 0 {
			if string(row[pos%width]) == "#" {
				count++
			}
			pos += step.across
		}
	}

	return count
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	trees := 1
	steps := [5]slope{slope{1, 1}, slope{3, 1}, slope{5, 1}, slope{7, 1}, slope{1, 2}}
	for _, step := range steps {

		trees *= calculateTrees(lines, step)
	}
	fmt.Println("count ", trees)
}
