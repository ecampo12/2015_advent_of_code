package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func findCombinations(targetSum int, nums []int, start int, currentSum int, currentCombo []int, result *[][]int) {
	if currentSum == targetSum {
		temp := make([]int, len(currentCombo))
		copy(temp, currentCombo)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		if currentSum+nums[i] <= targetSum {
			findCombinations(targetSum, nums, i+1, currentSum+nums[i], append(currentCombo, nums[i]), result)
		}
	}
}

func containerCombos(nums []int, targetSum int) [][]int {
	var result [][]int
	findCombinations(targetSum, nums, 0, 0, []int{}, &result)
	return result
}

func part1(input []string, target int) int {
	containers := []int{}
	for _, line := range input {
		x := 0
		fmt.Sscanf(line, "%d", &x)
		containers = append(containers, x)
	}

	combos := containerCombos(containers, target)
	return len(combos)
}

func part2(input []string, target int) int {
	containers := []int{}
	for _, line := range input {
		x := 0
		fmt.Sscanf(line, "%d", &x)
		containers = append(containers, x)
	}

	combos := containerCombos(containers, target)

	sort.Slice(combos, func(i, j int) bool {
		return len(combos[i]) < len(combos[j])
	})

	minLen := len(combos[0])
	count := 0

	for _, combo := range combos {
		if len(combo) != minLen {
			break
		}
		count++
	}

	return count
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines, 150)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines, 150)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
