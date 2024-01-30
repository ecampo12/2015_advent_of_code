package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"turn on 0,0 through 999,999"}, 1_000_000},
		{[]string{"toggle 0,0 through 999,0"}, 1000},
		{[]string{"turn off 499,499 through 500,500"}, 0},
		{
			[]string{
				"turn on 0,0 through 999,999",
				"toggle 0,0 through 999,0",
				"turn off 499,499 through 500,500",
			},
			998996,
		},
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
		{[]string{"turn on 0,0 through 0,0"}, 1},
		{[]string{"toggle 0,0 through 999,999"}, 2_000_000},
		{
			[]string{
				"turn on 0,0 through 0,0",
				"toggle 0,0 through 999,999",
			},
			2000001,
		},
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
