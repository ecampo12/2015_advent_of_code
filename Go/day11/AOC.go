package main

import (
	"fmt"
	"os"
	"strings"
)

// Password requirements:
// - It must contain at least one increasing straight of three letters. (e.g. abc, bcd, cde, etc.)
// - It cannot contain the letters i, o, or l.
// - It must contain at least two different, non-overlapping pairs of letters. (e.g. aa, bb, or zz)
func passwordValidator(password string) bool {
	// Check for increasing straight of three letters
	increasing := false
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			increasing = true
			break
		}
	}

	if !increasing {
		return false
	}

	// Check for i, o, or l
	if strings.Contains(password, "i") || strings.Contains(password, "o") || strings.Contains(password, "l") {
		return false
	}

	// Check for at least two different, non-overlapping pairs of letters
	pairs := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++
		}
	}
	return pairs >= 2
}

func nextPassword(password string) string {
	p := []byte(password)
	if p[len(p)-1] == 'z' {
		p[len(p)-1] = 'a'
		index := len(p) - 2
		for index >= 0 && p[index] == 'z' {
			p[index] = 'a'
			index--
		}
		if index >= 0 {
			p[index]++
		} else {
			p = append([]byte{'a'}, p...)
		}
	} else {
		p[len(p)-1]++
	}
	return string(p)
}

func nextValidPassword(password string) string {
	for {
		// fmt.Printf("Checking %s\n", password)
		if password == "ghjaabcc" {
			fmt.Print("Found it\n")
		}
		password = nextPassword(password)
		if passwordValidator(password) {
			return password
		}
	}
}

func part1(input string) string {
	return nextValidPassword(input)
}

func part2(input string) string {
	return nextValidPassword(input)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines[0])
	fmt.Printf("Part 1: %s\n", part1_ans)

	part2_ans := part2(part1_ans)
	fmt.Printf("Part 2: %s\n", part2_ans)
}
