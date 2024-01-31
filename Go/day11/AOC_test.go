package main

import (
	"testing"
)

func TestPasswordValidator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hijklmmn", false},
		{"abbceffg", false},
		{"abbcegjk", false},
	}

	for _, test := range tests {
		actual := passwordValidator(test.input)
		if actual != test.expected {
			t.Errorf("%s = %t, expected %t", test.input, actual, test.expected)
		}
	}
}

func TestNextPassword(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"xx", "xy"},
		{"yy", "yz"},
	}

	for _, test := range tests {
		actual := nextPassword(test.input)
		if actual != test.expected {
			t.Errorf("%s = %s, expected %s", test.input, actual, test.expected)
		}
	}
}

func TestNextValidPassword(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for _, test := range tests {
		actual := nextValidPassword(test.input)
		if actual != test.expected {
			t.Errorf("%s = %s, expected %s", test.input, actual, test.expected)
		}
	}
}
