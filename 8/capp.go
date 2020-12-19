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

func processLines(lines []string) (int, int) {
	var acc int
	var visited map[int]bool = make(map[int]bool)
	i := 0
	var j int
	for visited[i] != true {
		visited[i] = true
		j = i
		switch words := strings.Split(lines[i], " "); words[0] {
		case "nop":
			i++
		case "acc":
			val, _ := strconv.Atoi(words[1])
			acc += val
			i++
		case "jmp":
			val, _ := strconv.Atoi(words[1])
			i += val
		default:
			fmt.Println("incorrect input")
		}
	}
	return acc, j
}

func main() {
	lines, err := readLines("input")
	//lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	acc, i := processLines(lines)
	fmt.Println(acc, i)
}
