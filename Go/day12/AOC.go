package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input string) int {
	total := 0
	nums := regexp.MustCompile(`-?\d+`).FindAllString(input, -1)
	for _, num := range nums {
		x, _ := strconv.Atoi(num)
		total += x
	}
	return total
}

func getSum(data interface{}) int {
	total := 0
	switch t := data.(type) {
	case float64:
		total += int(t)
	case []interface{}:
		for _, v := range t {
			total += getSum(v)
		}
	case map[string]interface{}:
		for _, v := range t {
			if v == "red" {
				return 0
			}
			total += getSum(v)
		}
	}
	return total
}

func part2(input string) int {
	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		fmt.Print("Error unmarshaling json\n")
		panic(err)
	}

	return getSum(data)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines[0])
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines[0])
	fmt.Printf("Part 2: %d\n", part2_ans)
}
