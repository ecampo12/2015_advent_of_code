package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type location string
type distance int

type node struct {
	name      location
	neighbors map[location]distance
}

func (n *node) addNeighbor(loc location, dist distance) {
	if n.neighbors == nil {
		n.neighbors = make(map[location]distance)
	}
	n.neighbors[loc] = dist
}

func parseInput(input []string) map[location]*node {
	nodes := make(map[location]*node)

	for _, line := range input {
		var loc1, loc2 location
		var dist distance
		fmt.Sscanf(line, "%s to %s = %d", &loc1, &loc2, &dist)

		if _, ok := nodes[loc1]; !ok {
			nodes[loc1] = &node{name: loc1}
		}
		if _, ok := nodes[loc2]; !ok {
			nodes[loc2] = &node{name: loc2}
		}

		nodes[loc1].addNeighbor(loc2, dist)
		nodes[loc2].addNeighbor(loc1, dist)
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

func generateRoutes(nodes map[location]*node) []int {
	var routes []int
	locs := []string{}
	for loc := range nodes {
		locs = append(locs, string(loc))
	}

	perms := generatePermutations(locs)
	for _, perm := range perms {
		dist := 0
		for i := 0; i < len(perm)-1; i++ {
			dist += int(nodes[location(perm[i])].neighbors[location(perm[i+1])])
		}
		routes = append(routes, dist)
	}

	return routes
}

func part1(input []int) int {
	slices.Sort(input)
	return input[0]
}

func part2(input []int) int {
	slices.Sort(input)
	return input[len(input)-1]
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	routes := generateRoutes(parseInput(lines))

	part1_ans := part1(routes)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(routes)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
