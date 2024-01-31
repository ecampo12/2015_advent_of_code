package main

import (
	"slices"
	"testing"
)

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

func TestGenerateRoutes(t *testing.T) {
	tests := []struct {
		input    []string
		expected []int
	}{
		{
			[]string{
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			},
			[]int{605, 605, 659, 659, 982, 982},
		},
	}

	for _, test := range tests {
		actual := generateRoutes(parseInput(test.input))
		slices.Sort(actual)
		if !compareSlices(actual, test.expected) {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			},
			605,
		},
	}

	for _, test := range tests {
		routes := generateRoutes(parseInput(test.input))
		actual := part1(routes)
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
				"London to Dublin = 464",
				"London to Belfast = 518",
				"Dublin to Belfast = 141",
			},
			982,
		},
	}

	for _, test := range tests {
		routes := generateRoutes(parseInput(test.input))
		actual := part2(routes)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
