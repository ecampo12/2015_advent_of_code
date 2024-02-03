package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Sue struct {
	id         int
	properties map[string]int
}

func parseInput(input []string) []Sue {
	sue := []Sue{}

	for i, line := range input {
		props := regexp.MustCompile(`(\w+): (\d+)`).FindAllString(line, -1)
		s := Sue{i + 1, make(map[string]int)}

		for _, p := range props {
			data := strings.Split(p, ": ")
			x, _ := strconv.Atoi(data[1])
			s.properties[data[0]] = x
		}

		sue = append(sue, s)
	}

	return sue
}

func part1(input []string) int {
	sue := parseInput(input)
	search := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	for _, s := range sue {
		match := true
		for k, v := range s.properties {
			if search[k] != v {
				match = false
				break
			}
		}
		if match {
			return s.id
		}
	}

	return 0
}

func part2(input []string) int {
	sue := parseInput(input)

	search := map[string]int{
		"children": 3,
		"samoyeds": 2,
		"akitas":   0,
		"vizslas":  0,
		"cars":     2,
		"perfumes": 1,
	}

	for _, s := range sue {
		match := true
		for k, v := range s.properties {
			if k == "cats" || k == "trees" {
				if search[k] >= v {
					match = false
					break
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if search[k] <= v {
					match = false
					break
				}
			} else if search[k] != v {
				match = false
				break
			}
		}
		if match {
			return s.id
		}
	}

	return 0
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
