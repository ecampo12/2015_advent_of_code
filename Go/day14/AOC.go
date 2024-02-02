package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Reindeer struct {
	name     string
	speed    int
	flyTime  int
	restTime int
}

func parseInput(input []string) []Reindeer {
	reindeer := []Reindeer{}

	var name string
	var speed, flyTime, restTime int
	for _, line := range input {
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &flyTime, &restTime)
		reindeer = append(reindeer, Reindeer{name, speed, flyTime, restTime})
	}

	return reindeer
}

func part1(input []string, racetime int) int {
	r := parseInput(input)
	var distance []int

	for _, deer := range r {
		cycleTime := deer.flyTime + deer.restTime
		cycles := racetime / cycleTime
		remainingTime := racetime % cycleTime
		distance = append(distance, (cycles*deer.flyTime+min(deer.flyTime, remainingTime))*deer.speed)
	}

	slices.Sort(distance)
	return distance[len(distance)-1]
}

func part2(input []string, racetime int) int {
	r := parseInput(input)
	points := make([]int, len(r))
	distance := make([]int, len(r))

	for i := 1; i <= racetime; i++ {
		furthest := 0
		for j, deer := range r {
			cycleTime := deer.flyTime + deer.restTime
			cycles := i / cycleTime
			remainingTime := i % cycleTime
			distance[j] = (cycles*deer.flyTime + min(deer.flyTime, remainingTime)) * deer.speed
			if distance[j] > furthest {
				furthest = distance[j]
			}
		}

		for j, d := range distance {
			if d == furthest {
				points[j]++
			}
		}
	}

	slices.Sort(points)
	return points[len(points)-1]
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines, 2503)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines, 2503)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
