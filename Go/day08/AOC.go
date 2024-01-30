package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func part1(input []string) int {
	hex := regexp.MustCompile(`\\x[0-9a-fA-F]{2}`)
	total := 0

	for _, line := range input {
		numOfChars := len(line)
		line = line[1 : len(line)-1]
		line = strings.ReplaceAll(line, `\"`, `"`)
		line = hex.ReplaceAllString(line, `X`)
		line = strings.ReplaceAll(line, `\\`, `X`)
		total += numOfChars - len(line)
	}
	return total
}

func part2(input []string) int {
	total := 0
	for _, line := range input {
		total += strings.Count(line, `"`) + strings.Count(line, `\`) + 2
	}
	return total
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
