package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{[]string{"[1,2,3]"}, 6},
		{[]string{"{\"a\":2,\"b\":4}"}, 6},
		{[]string{"[[[3]]]"}, 3},
		{[]string{"{\"a\":{\"b\":4},\"c\":-1}"}, 3},
		{[]string{"{\"a\":[-1,1]}"}, 0},
		{[]string{"[-1,{\"a\":1}]"}, 0},
		{[]string{"[]"}, 0},
		{[]string{"{}"}, 0},
	}

	for _, test := range tests {
		actual := part1(test.input[0])
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
		{[]string{"[1,2,3]"}, 6},
		{[]string{"[1,{\"c\":\"red\",\"b\":2},3]"}, 4},
		{[]string{"{\"d\":\"red\",\"e\":[1,2,3,4],\"f\":5}"}, 0},
		{[]string{"[1,\"red\",5]"}, 6},
	}

	for _, test := range tests {
		actual := part2(test.input[0])
		if actual != test.expected {
			t.Errorf("%s = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
