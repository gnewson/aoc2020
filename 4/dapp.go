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

func validateYear(year string, lower int, upper int) bool {
	var ret bool

	//fmt.Printf("\"%s\"\n", year)

	if matched, _ := regexp.MatchString(`^\d{4}$`, year); matched == true {
		if byr, err := strconv.Atoi(year); err == nil {
			if byr >= lower && byr <= upper {
				ret = true
			} else {
				ret = false
			}
		} else {
			ret = false
		}

	} else {
		ret = false
	}

	return ret
}

func validatePid(pid string) bool {
	var ret bool
	if len(pid) == 9 {
		if matched, _ := regexp.MatchString(`^0*\d+$`, pid); matched == true {
			ret = true
		} else {
			ret = false
		}
	} else {
		ret = false
	}
	return ret
}

func validateEcl(ecl string) bool {
	var ret bool
	if matched, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, ecl); matched == true {
		ret = true
	} else {
		ret = false
	}
	return ret
}

func validateHcl(hcl string) bool {
	var ret bool
	if matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, hcl); matched == true {
		ret = true
	} else {
		ret = false
	}
	return ret
}

func validateHgt(hgt string) bool {
	var ret bool

	var re = regexp.MustCompile(`(?m)^(\d\d)in$|^(1\d\d)cm$`)
	var res = re.FindStringSubmatch(hgt)

	if len(res) != 0 {
		if res[1] != "" {
			value, _ := strconv.Atoi(res[1])
			if value >= 59 && value <= 76 {
				ret = true
			} else {
				ret = false
			}
		} else {
			value, _ := strconv.Atoi(res[2])
			if value >= 150 && value <= 193 {
				ret = true
			} else {
				ret = false
			}
		}

	} else {
		ret = false
	}
	return ret
}

func validId(id record) bool {
	var ret bool = true
	if !validateYear(id.eyr, 2020, 2030) {
		ret = false
	}

	if !validateYear(id.iyr, 2010, 2020) {
		ret = false
	}

	if !validateYear(id.byr, 1920, 2002) {
		ret = false
	}

	if !validatePid(id.pid) {
		ret = false
	}

	if !validateEcl(id.ecl) {
		ret = false
	}

	if !validateHcl(id.hcl) {
		ret = false
	}

	if !validateHgt(id.hgt) {
		ret = false
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
