package main

import (
	"testing"
)

func TestDay10(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 2},
		{"11", 2},
		{"21", 4},
		{"1211", 6},
		{"111221", 6},
	}

	for _, test := range tests {
		actual := len(lookAndSay(test.input))
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
