package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(input []string) [][]string {
	grid := [][]string{}

	for _, line := range input {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func neighborsOn(grid [][]string, x, y int) int {
	neighbors := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if (i == 0 && j == 0) || x+i < 0 || y+j < 0 || x+i >= len(grid) || y+j >= len(grid[0]) {
				continue
			}
			if grid[x+i][y+j] == "#" {
				neighbors++
			}
		}
	}
	return neighbors
}

func turnOn(grid [][]string, alwaysOn bool) [][]string {
	if alwaysOn {
		grid[0][0] = "#"
		grid[0][len(grid[0])-1] = "#"
		grid[len(grid)-1][0] = "#"
		grid[len(grid)-1][len(grid[0])-1] = "#"
	}

	newGrid := make([][]string, len(grid))
	for i := range grid {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "#" {
				if neighborsOn(grid, i, j) != 2 && neighborsOn(grid, i, j) != 3 {
					newGrid[i][j] = "."
				}
			} else {
				if neighborsOn(grid, i, j) == 3 {
					newGrid[i][j] = "#"
				}
			}
		}
	}

	return newGrid
}

func part1(input []string, steps int, alwaysOn bool) int {
	grid := parseInput(input)

	for i := 0; i < steps; i++ {
		grid = turnOn(grid, alwaysOn)

		if alwaysOn {
			grid[0][0] = "#"
			grid[0][len(grid[0])-1] = "#"
			grid[len(grid)-1][0] = "#"
			grid[len(grid)-1][len(grid[0])-1] = "#"
		}
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == "#" {
				count++
			}
		}
	}
	return count
}

func part2(input []string, steps int) int {
	return part1(input, steps, true)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines, 100, false)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines, 100)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
