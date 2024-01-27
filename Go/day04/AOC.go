package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func part1(input []byte) int {
	count := 0
	for {
		i := []byte(fmt.Sprintf("%d", count))
		str := append(input, i...)
		hash := md5.Sum(str)
		if hex.EncodeToString(hash[:])[:5] == "00000" {
			return count
		}
		count++
	}
}

func part2(input []byte) int {
	count := 0
	for {
		i := []byte(fmt.Sprintf("%d", count))
		str := append(input, i...)
		hash := md5.Sum(str)
		if hex.EncodeToString(hash[:])[:6] == "000000" {
			return count
		}
		count++
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	start := time.Now()
	part1_ans := part1(input)
	duration := time.Since(start)
	fmt.Printf("Part 1: %d\n", part1_ans)
	fmt.Printf("Part 1 took %s\n\n", duration)

	start = time.Now()
	part2_ans := part2(input)
	duration = time.Since(start)
	fmt.Printf("Part 2: %d\n", part2_ans)
	fmt.Printf("Part 2 took %s\n", duration)
}
