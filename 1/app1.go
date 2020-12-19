package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		linevalue, _ := strconv.Atoi(line)
		for _, jline := range lines[i+1:] {
			jvalue, _ := strconv.Atoi(jline)
			if linevalue+jvalue == 2020 {
				fmt.Printf("%d %d\n", linevalue, jvalue)
				fmt.Printf("%d\n", linevalue*jvalue)
			}
		}
	}
}
