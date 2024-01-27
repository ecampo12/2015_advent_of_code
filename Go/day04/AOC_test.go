package main

import (
	"testing"
)

func TestDay4(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, test := range tests {
		actual := part1([]byte(test.input))
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
