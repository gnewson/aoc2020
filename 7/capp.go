package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func processLines(lines []string, colour string) int {
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)
	var rowCount int
	for _, line := range lines {
		words := strings.Split(line, " ")
		mainColour := words[0] + " " + words[1]
		if mainColour == colour {
			for i, w := range words {
				if digitCheck.MatchString(w) {
					num, _ := strconv.Atoi(w)
					rowCount += num + num*processLines(lines, words[i+1]+" "+words[i+2])
				}
			}
		}
	}
	return rowCount
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	bags := processLines(lines, "shiny gold")
	fmt.Println(bags)
}
