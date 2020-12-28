package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}
	slice := lines[:]
	sort.Ints(slice)
	return slice, scanner.Err()
}

func findJolts(lines []int) (int, int) {
	previous := 0
	var threes int
	var ones int
	for i, current := range lines {
		if current-previous == 1 {
			ones++
		} else if current-previous == 3 {
			threes++
		}
		previous = lines[i]
	}
	threes++
	return ones, threes
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	ones, threes := findJolts(lines)
	fmt.Println("Number:", ones*threes)
}
