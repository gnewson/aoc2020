package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var PREAMBLE int = 25

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

	return lines, scanner.Err()
}

func findNumber(lines []int) int {
	var output int
	for i, number := range lines[PREAMBLE:] {
		var sums map[int]int = make(map[int]int)

		for j, jval := range lines[i : i+PREAMBLE] {
			for _, kval := range lines[j+1 : i+PREAMBLE] {
				sums[jval+kval] += 1
			}
		}

		if _, in := sums[number]; in == false {
			output = number
			break
		}
	}
	return output
}

func findContiguousSum(lines []int, number int) int {
	fmt.Println("number:", number)
	var runningTot int
	var returnSum int
	exit := false
	for i, _ := range lines[PREAMBLE:] {
		fmt.Println(i)
		if exit {
			break
		}
		for j, value := range lines[PREAMBLE+i:] {
			runningTot += value
			if runningTot == number {
				//find min(lines[:i]
				//find max(lines[:i]
				//sum min and max
				fmt.Println(minMax(lines[PREAMBLE+i : PREAMBLE+i+j+1]))
				min, max := minMax(lines[PREAMBLE+i : PREAMBLE+i+j+1])
				returnSum = min + max
				exit = true
				break
			} else if runningTot > number {
				runningTot = 0
				break
			}
		}
	}
	return returnSum
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println("sum is:", findContiguousSum(lines, findNumber(lines)))
}
