package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	name      string
	neighbors map[string]int
}

func (n *node) addNeighbor(person string, happiness int) {
	if n.neighbors == nil {
		n.neighbors = make(map[string]int)
	}
	n.neighbors[person] = happiness
}

func parseInput(input []string) map[string]*node {
	nodes := make(map[string]*node)

	for _, line := range input {
		var name1, name2 string
		var op string
		var happiness int
		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s.", &name1, &op, &happiness, &name2)

		if op == "lose" {
			happiness = -happiness
		}

		if name2[len(name2)-1] == '.' {
			name2 = name2[:len(name2)-1]
		}

		if _, ok := nodes[name1]; !ok {

			nodes[name1] = &node{name: name1}
		}
		if _, ok := nodes[name2]; !ok {
			nodes[name2] = &node{name: name2}
		}

		nodes[name1].addNeighbor(name2, happiness)
	}

	return nodes
}

func generatePermutations(elements []string) [][]string {
	var result [][]string

	var generate func(int)
	generate = func(n int) {
		if n == 1 {
			temp := make([]string, len(elements))
			copy(temp, elements)
			result = append(result, temp)
			return
		}

		for i := 0; i < n; i++ {
			elements[i], elements[n-1] = elements[n-1], elements[i]
			generate(n - 1)
			elements[i], elements[n-1] = elements[n-1], elements[i]
		}
	}

	generate(len(elements))
	return result
}

func generateHappiness(nodes map[string]*node) int {
	var max int
	people := make([]string, 0, len(nodes))
	for name := range nodes {
		people = append(people, name)
	}

	permutations := generatePermutations(people)

	for _, perm := range permutations {
		mood := 0
		for i := 0; i < len(perm); i++ {
			mood += nodes[perm[i]].neighbors[perm[(i+1)%len(perm)]]
			if i == 0 {
				mood += nodes[perm[i]].neighbors[perm[len(perm)-1]]
			} else {
				mood += nodes[perm[i]].neighbors[perm[(i-1)]]
			}
		}
		if mood > max {
			max = mood
		}
	}

	return max
}

func part1(input []string) int {
	return generateHappiness(parseInput(input))
}

func part2(input []string) int {
	nodes := parseInput(input)
	nodes["me"] = &node{name: "me"}
	for name := range nodes {
		nodes[name].addNeighbor("me", 0)
		nodes["me"].addNeighbor(name, 0)
	}

	return generateHappiness(nodes)
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
