package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"20",
				"15",
				"10",
				"5",
				"5",
			},
			4,
		},
	}

	for _, test := range tests {
		actual := part1(test.input, 25)
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
		{
			[]string{
				"20",
				"15",
				"10",
				"5",
				"5",
			},
			3,
		},
	}

	for _, test := range tests {
		actual := part2(test.input, 25)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
