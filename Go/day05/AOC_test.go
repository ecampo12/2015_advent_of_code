package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"ugknbfddgicrmopn"}, 1},
		{[]string{"aaa"}, 1},
		{[]string{"jchzalrnumimnmhp"}, 0},
		{[]string{"haegwjzuvuyypxyu"}, 0},
		{[]string{"dvszwmarrgswjxmb"}, 0},
	}

	for _, test := range tests {
		actual := part1(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"qjhvhtzxzqqjkmpb"}, 1},
		{[]string{"xxyxx"}, 1},
		{[]string{"uurcxstgmygtbstg"}, 0},
		{[]string{"ieodomkazucvgmuy"}, 0},
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
