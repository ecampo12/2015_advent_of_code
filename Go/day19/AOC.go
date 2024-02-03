package main

import (
	"fmt"
	"os"
	"strings"
)

func part1(input []string) int {
	start := input[len(input)-1]
	input = input[:len(input)-2]
	molecules := make(map[string]bool)

	for _, line := range input {
		parts := strings.Split(line, " => ")
		for i := 0; i < len(start); i++ {
			if i+len(parts[0]) <= len(start) && start[i:i+len(parts[0])] == parts[0] {
				molecules[start[:i]+parts[1]+start[i+len(parts[0]):]] = true
			}
		}
	}
	return len(molecules)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// I think I figured out why the test cases in Python soltuion keept failing.
// I do not think the problem was design to be solved in reverse. Because our
// approach results in the test cases producing strings of 'e's instaed of
// one 'e' as the expected output. The input works out fine.
func part2(input []string) int {
	start := reverse(input[len(input)-1])
	input = input[:len(input)-2]

	replacements := map[string]string{}
	for _, line := range input {
		parts := strings.Split(line, " => ")
		replacements[reverse(parts[1])] = reverse(parts[0])
	}

	steps := 0

	for !strings.Contains(start, "e") {
		for i := 0; i < len(start); i++ {
			for j := 1; j <= len(start); j++ {
				if j+i <= len(start) {
					if _, ok := replacements[start[i:j+i]]; ok {
						start = start[:i] + replacements[start[i:j+i]] + start[j+i:]
						steps++
						break
					}
				}
			}
		}
	}
	return steps
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
