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
				"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
				"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			},
			62842880,
		},
	}

	for _, test := range tests {
		actual := part1(test.input, false)
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
				"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
				"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			},
			57600000,
		},
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
