package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"\"\""}, 2},
		{[]string{"\"abc\""}, 2},
		{[]string{"\"aaa\\\"aaa\""}, 3},
		{[]string{"\"\\x27\""}, 5},
		{
			[]string{
				"\"\"",
				"\"abc\"",
				"\"aaa\\\"aaa\"",
				"\"\\x27\"",
			},
			12,
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
		{[]string{"\"\""}, 4},
		{[]string{"\"abc\""}, 4},
		{[]string{"\"aaa\\\"aaa\""}, 6},
		{[]string{"\"\\x27\""}, 5},
		{
			[]string{
				"\"\"",
				"\"abc\"",
				"\"aaa\\\"aaa\"",
				"\"\\x27\"",
			},
			19,
		},
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
