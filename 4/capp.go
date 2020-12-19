package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type record struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
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
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newline := scanner.Text()
		if newline == "" {
			lines = append(lines, strings.TrimSpace(line))
			line = ""
		} else {
			line += " " + newline
		}
	}

	lines = append(lines, strings.TrimSpace(line))
	return lines, scanner.Err()
}

func createStruct(line string) record {
	fields := strings.Split(line, " ")

	var id record

	for _, field := range fields {

		switch parts := strings.Split(field, ":"); parts[0] {
		case "eyr":
			id.eyr = parts[1]
		case "hgt":
			id.hgt = parts[1]
		case "hcl":
			id.hcl = parts[1]
		case "ecl":
			id.ecl = parts[1]
		case "pid":
			id.pid = parts[1]
		case "cid":
			id.cid = parts[1]
		case "byr":
			id.byr = parts[1]
		case "iyr":
			id.iyr = parts[1]
		default:
			// freebsd, openbsd,
			// plan9, windows...

			fmt.Println(id)
			fmt.Printf("%s.\n", parts[0])
		}
	}

	return id
}

func validId(id record) bool {
	var ret bool = false
	if id.eyr != "" &&
		id.iyr != "" &&
		id.byr != "" &&
		id.pid != "" &&
		id.ecl != "" &&
		id.hcl != "" &&
		id.hgt != "" {
		ret = true
	}

	return ret
}

func main() {
	lines, err := readLines("input")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	records := []record{}
	for _, line := range lines {
		rec := createStruct(line)
		records = append(records, rec)
	}

	var count int
	for _, id := range records {
		if validId(id) {
			count++
		}
	}

	fmt.Println("Number of valid id: ", count)

}
