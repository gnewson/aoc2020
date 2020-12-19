package main

import (
	"fmt"
	"testing"
)

func TestValidateYear(t *testing.T) {
	tests := []struct {
		year  string
		lower int
		upper int
		valid bool
	}{
		{"2002", 1920, 2002, true},
		{"2003", 1920, 2002, false},
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

func TestValidatePid(t *testing.T) {
	tests := []struct {
		pid   string
		valid bool
	}{
		{"000000001", true},
		{"0123456789", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.pid)
		t.Run(testname, func(t *testing.T) {
			ans := validatePid(tt.pid)
			if ans != tt.valid {
				t.Errorf("got %t, want %t", ans, tt.valid)
			}
		})
	}
}

func TestValidateEcl(t *testing.T) {
	tests := []struct {
		colour string
		valid  bool
	}{
		{"brn", true},
		{"wat", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.colour)
		t.Run(testname, func(t *testing.T) {
			ans := validateEcl(tt.colour)
			if ans != tt.valid {
				t.Errorf("got %t, want %t", ans, tt.valid)
			}
		})
	}
}

func TestValidateHcl(t *testing.T) {
	tests := []struct {
		colour string
		valid  bool
	}{
		{"#123abc", true},
		{"#123abz", false},
		{"123abc", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.colour)
		t.Run(testname, func(t *testing.T) {
			ans := validateHcl(tt.colour)
			if ans != tt.valid {
				t.Errorf("got %t, want %t", ans, tt.valid)
			}
		})
	}
}

func TestValidateHgt(t *testing.T) {
	tests := []struct {
		hgt   string
		valid bool
	}{
		{"60in", true},
		{"190cm", true},
		{"190in", false},
		{"60", false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.hgt)
		t.Run(testname, func(t *testing.T) {
			ans := validateHgt(tt.hgt)
			if ans != tt.valid {
				t.Errorf("got %t, want %t", ans, tt.valid)
			}
		})
	}
}
