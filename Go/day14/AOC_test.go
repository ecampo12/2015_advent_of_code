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
				"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
				"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
			},
			1120,
		},
	}

	for _, test := range tests {
		actual := part1(test.input, 1000)
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
				"Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.",
				"Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.",
			},
			689,
		},
	}

	for _, test := range tests {
		actual := part2(test.input, 1000)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
