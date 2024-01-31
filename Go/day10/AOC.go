package main

import (
	"fmt"
	"os"
	"strings"
)

func lookAndSay(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			i++
			count++
		}
		result.WriteString(fmt.Sprintf("%d%c", count, s[i]))
	}
	return result.String()
}

func part1(input []string) int {
	init := input[0]
	for i := 0; i < 40; i++ {
		init = lookAndSay(init)
	}
	return len(init)
}

func part2(input []string) int {
	init := input[0]
	for i := 0; i < 50; i++ {
		init = lookAndSay(init)
	}
	return len(init)
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
