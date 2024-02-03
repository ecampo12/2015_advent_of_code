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
				"H => HO",
				"H => OH",
				"O => HH",
				"",
				"HOH",
			},
			4,
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
		{
			[]string{
				"e => H",
				"e => O",
				"H => HO",
				"H => OH",
				"O => HH",
				"",
				"HOH",
			},
			3,
		},
		{
			[]string{
				"e => H",
				"e => O",
				"H => HO",
				"H => OH",
				"O => HH",
				"",
				"HOHOHO",
			},
			6,
		},
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
