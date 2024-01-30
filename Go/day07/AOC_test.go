package main

import (
	"testing"
)

func TestDay07(t *testing.T) {
	input := []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
		"x -> a",
	}

	expected := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
		"a": 123,
	}

	actual := part1(input)

	for k, v := range expected {
		if actual[k] != v {
			t.Errorf("%s = %d, expected %d", k, actual[k], v)
		}
	}
}
