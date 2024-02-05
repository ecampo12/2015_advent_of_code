package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var cache map[int][]int

func getFactors(n int) []int {
	factors := []int{}

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			factors = append(factors, i)

			if i != n/i {
				factors = append(factors, n/i)
			}
		}
	}
	cache[n] = factors
	return factors
}

func part1(input []string) int {
	target, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Print("Error converting to int\n")
		panic(err)
	}

	for i := 1; i < target; i++ {
		presents := 0
		f := getFactors(i)
		for _, factor := range f {
			presents += factor * 10
		}
		if presents >= target {
			return i
		}
	}
	return 0
}

func part2(input []string) int {
	target, err := strconv.Atoi(input[0])
	if err != nil {
		fmt.Print("Error converting to int\n")
		panic(err)
	}

	for i := 1; i < target; i++ {
		presents := 0
		var f []int
		if _, ok := cache[i]; ok {
			f = cache[i]
		} else {
			f = getFactors(i)
		}
		for _, factor := range f {
			if i/factor <= 50 {
				presents += factor * 11
			}
		}
		if presents >= target {
			return i
		}
	}

	return 0
}

func main() {
	cache = make(map[int][]int)
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	start := time.Now()
	part1_ans := part1(lines)
	elapsed1 := time.Since(start)
	fmt.Printf("Part 1: %d\n", part1_ans)

	start = time.Now()
	part2_ans := part2(lines)
	elapsed2 := time.Since(start)
	fmt.Printf("Part 2: %d\n\n", part2_ans)

	fmt.Printf("Part 1 took %s\n", elapsed1)
	fmt.Printf("Part 2 took %s\n", elapsed2)
}
