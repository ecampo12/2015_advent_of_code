package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func part1(input []string) int {
	nice := 0
	for _, line := range input {
		vowels := regexp.MustCompile(`[aeiou]`)
		badStrings := regexp.MustCompile(`ab|cd|pq|xy`)

		if len(vowels.FindAllString(line, -1)) >= 3 && hasDoubleLetter(line) && !badStrings.MatchString(line) {
			nice++
		}

	}
	return nice
}

// Ran into trouble getting backreferences to work in Go, so I just used a loop
func hasDoubleLetter(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func part2(input []string) int {
	nice := 0
	for _, line := range input {
		if hasDoublePair(line) && hasSandwich(line) {
			nice++
		}
	}
	return nice
}

func hasDoublePair(line string) bool {
	for i := 0; i < len(line)-1; i++ {
		if strings.Contains(line[i+2:], line[i:i+2]) {
			return true
		}
	}
	return false
}

func hasSandwich(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
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
