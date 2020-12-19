package main

import (
	"fmt"
	"testing"
)

func TestCreateStruct(t *testing.T) {
	tests := []struct {
		details       string
		structDetails record
	}{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
			record{"1937", "2017", "2020", "183cm", "#fffffd", "gry", "860033327", "147"}},
		{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
			record{"1929", "2013", "2023", "", "#cfa07d", "amb", "028048884", "350"}},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
			record{"1931", "2013", "2024", "179cm", "#ae17e1", "brn", "760753108", ""}},
		{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
			record{"", "2011", "2025", "59in", "#cfa07d", "brn", "166559648", ""}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.details)
		t.Run(testname, func(t *testing.T) {
			ans := createStruct(tt.details)
			if ans != tt.structDetails {
				t.Errorf("got %s, want %s", ans, tt.structDetails)
			}
		})
	}
}

func TestValidId(t *testing.T) {
	tests := []struct {
		structDetails record
		valid         bool
	}{
		{record{"1937", "2017", "2020", "183cm", "#fffffd", "gry", "860033327", "147"}, true},
		{record{"1929", "2013", "2023", "", "#cfa07d", "amb", "028048884", "350"}, false},
		{record{"1931", "2013", "2024", "179cm", "#ae17e1", "brn", "760753108", ""}, true},
		{record{"", "2011", "2025", "59in", "#cfa07d", "brn", "166559648", ""}, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.structDetails)
		t.Run(testname, func(t *testing.T) {
			ans := validId(tt.structDetails)
			if ans != tt.valid {
				t.Errorf("got %t, want %t", ans, tt.valid)
			}
		})
	}
}

func TestValidateYear(t *testing.T) {
	tests := []struct {
		year  string
		lower int
		upper int
		valid bool
	}{
		{"2002", 1920, 2002, true},
		{"2003", 1920, 2002, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.year)
		t.Run(testname, func(t *testing.T) {
			ans := validateYear(tt.year, tt.lower, tt.upper)
			if ans != tt.valid {
				t.Errorf("got %t, want %t", ans, tt.valid)
			}
		})
	}
}
