package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

func processLines(lines []string) map[string][]string {
	var bags map[string][]string = make(map[string][]string)
	var digitCheck = regexp.MustCompile(`^[0-9]+$`)
	for _, line := range lines {
		words := strings.Split(line, " ")
		key := words[0] + " " + words[1]
		bags[key] = nil
		for i, w := range words {
			if digitCheck.MatchString(w) {

				bags[key] = append(bags[key], words[i+1]+" "+words[i+2])
			}
		}

	}
	//fmt.Println(bags)
	return bags
}

func find(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}

func unique(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func getBagColours(bags map[string][]string, colour string) []string {
	var colours []string
	for k, v := range bags {
		if find(v, colour) {
			colours = append(colours, k)
			colours = append(colours, getBagColours(bags, k)...)
			//return getBagColours(bags, k)
		}
	}
	//fmt.Println(colours)
	return unique(colours)
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

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	bags := processLines(lines)
	colours := getBagColours(bags, "shiny gold")
	fmt.Println(len(colours))
}
