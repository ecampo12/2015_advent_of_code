package main

import (
	"fmt"
	"os"
)

func part1(input []byte) int {
	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
	}
	return floor
}

func part2(input []byte) int {
	floor := 0
	for i, c := range input {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}
	return floor
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	part1_ans := part1(input)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(input)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
