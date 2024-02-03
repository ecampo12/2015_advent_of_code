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
				".#.#.#",
				"...##.",
				"#....#",
				"..#...",
				"#.#..#",
				"####..",
			},
			4,
		},
	}

	for _, test := range tests {
		actual := part1(test.input, 4, false)
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
				"##.#.#",
				"...##.",
				"#....#",
				"..#...",
				"#.#..#",
				"####.#",
			},
			17,
		},
	}

	for _, test := range tests {
		actual := part2(test.input, 5)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
