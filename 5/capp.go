package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

func getRow(fb string, lower int, upper int) int {
	midpoint := (upper + lower) / 2
	var value int

	if len(fb) == 1 {
		if fb == "F" {
			value = lower
		} else {
			value = upper
		}
	} else {
		if string(fb[0]) == "F" {
			return getRow(fb[1:], lower, midpoint)
		} else if string(fb[0]) == "B" {
			return getRow(fb[1:], midpoint+1, upper)
		}
	}

	return value
}

func getColumn(rl string, lower int, upper int) int {
	midpoint := (upper + lower) / 2
	var value int

	if len(rl) == 1 {
		if rl == "L" {
			value = lower
		} else {
			value = upper
		}
	} else {
		if string(rl[0]) == "L" {
			return getColumn(rl[1:], lower, midpoint)
		} else if string(rl[0]) == "R" {
			return getColumn(rl[1:], midpoint+1, upper)
		}
	}

	return value
}

func getSeatId(row int, column int) int {
	return row*8 + column
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var seatIds []int
	for _, line := range lines {
		row := getRow(line[:7], 0, 127)
		column := getColumn(line[7:], 0, 7)
		id := getSeatId(row, column)

		seatIds = append(seatIds, id)
	}

	sort.Ints(seatIds)

	var minusOne int

	for _, seatId := range seatIds {
		if seatId-minusOne == 2 {
			fmt.Println("Your seat is: ", seatId-1)
		} else {
			minusOne = seatId
		}
	}
}
