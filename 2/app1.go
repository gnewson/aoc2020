package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func parseLines(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		if validPassword(scanner.Text()) {
			count++
		}
	}
	fmt.Println(count)
	return scanner.Err()
}

func validPassword(line string) bool {
	tokens := strings.Fields(line)

	letter := strings.TrimSuffix(tokens[1], ":")
	pw := tokens[2]
	numbers := strings.Split(tokens[0], "-")
	lower, _ := strconv.Atoi(numbers[0])
	upper, _ := strconv.Atoi(numbers[1])

	count := strings.Count(pw, letter)

	if count <= upper && count >= lower {
		return true
	} else {
		return false
	}
}

func main() {
	err := parseLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
}
