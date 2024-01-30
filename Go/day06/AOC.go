package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) int {
	lights := make([][]bool, 1000)
	total := 0

	for i := range lights {
		lights[i] = make([]bool, 1000)
	}

	for _, line := range input {
		action := regexp.MustCompile(`turn on|turn off|toggle`).FindAllString(line, -1)[0]
		nums := regexp.MustCompile(`\d+`).FindAllString(line, -1)

		coords := make([]int, 4)
		for i, num := range nums {
			coords[i], _ = strconv.Atoi(num)
		}

		for x := coords[0]; x <= coords[2]; x++ {
			for y := coords[1]; y <= coords[3]; y++ {
				switch action {
				case "turn on":
					lights[x][y] = true
				case "turn off":
					lights[x][y] = false
				case "toggle":
					lights[x][y] = !lights[x][y]
				}
			}
		}
	}

	for _, row := range lights {
		for _, light := range row {
			if light {
				total++
			}
		}
	}

	return total
}

func part2(input []string) int {
	lights := make([][]int, 1000)
	total := 0

	for i := range lights {
		lights[i] = make([]int, 1000)
	}

	for _, line := range input {
		action := regexp.MustCompile(`turn on|turn off|toggle`).FindAllString(line, -1)[0]
		nums := regexp.MustCompile(`\d+`).FindAllString(line, -1)

		coords := make([]int, 4)
		for i, num := range nums {
			coords[i], _ = strconv.Atoi(num)
		}

		for x := coords[0]; x <= coords[2]; x++ {
			for y := coords[1]; y <= coords[3]; y++ {
				switch action {
				case "turn on":
					lights[x][y]++
				case "turn off":
					if lights[x][y] > 0 {
						lights[x][y]--
					}
				case "toggle":
					lights[x][y] += 2
				}
			}
		}
	}

	for _, row := range lights {
		for _, light := range row {
			total += light
		}
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
