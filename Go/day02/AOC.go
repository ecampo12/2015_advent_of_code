package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(input []string) int {
	total := 0
	for _, line := range input {
		var dims []int
		for _, dim := range strings.Split(line, "x") {
			d, _ := strconv.Atoi(dim)
			dims = append(dims, d)
		}

		l, w, h := dims[0], dims[1], dims[2]
		sides := []int{l * w, w * h, h * l}
		sort.Ints(sides)
		for _, side := range sides {
			total += 2 * side
		}

		total += sides[0]
	}
	return total
}

func part2(input []string) int {
	total := 0
	for _, line := range input {
		var dims []int
		for _, dim := range strings.Split(line, "x") {
			d, _ := strconv.Atoi(dim)
			dims = append(dims, d)
		}

		sort.Ints(dims)
		r1, r2 := dims[0], dims[1]

		total += 2*r1 + 2*r2 + r1*r2*dims[2]
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
