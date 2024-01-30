package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(input []string) map[string]uint16 {
	wires := make(map[string]uint16)

	for {
		wait := []string{}
		for _, line := range input {
			parts := strings.Split(line, " ")
			if regexp.MustCompile(`NOT`).FindAllString(line, -1) != nil {
				src := parts[1]
				wire := parts[3]
				if _, ok := wires[src]; ok {
					wires[wire] = ^wires[src]
				} else {
					wait = append(wait, line)
				}
			} else if regexp.MustCompile(`AND|OR|LSHIFT|RSHIFT`).FindAllString(line, -1) != nil {
				src1 := parts[0]
				op := parts[1]
				src2 := parts[2]
				wire := parts[4]
				num1, err1 := strconv.Atoi(src1)
				num2, err2 := strconv.Atoi(src2)

				val1, ok1 := wires[src1]
				val2, ok2 := wires[src2]

				if err1 != nil && !ok1 || err2 != nil && !ok2 {
					wait = append(wait, line)
				} else {
					var v1, v2 uint16
					if ok1 {
						v1 = val1
					} else {
						v1 = uint16(num1)
					}

					if ok2 {
						v2 = val2
					} else {
						v2 = uint16(num2)
					}

					switch op {
					case "AND":
						wires[wire] = v1 & v2
					case "OR":
						wires[wire] = v1 | v2
					case "LSHIFT":
						wires[wire] = v1 << v2
					case "RSHIFT":
						wires[wire] = v1 >> v2
					}
				}

			} else {
				src := parts[0]
				wire := parts[2]

				num, err := strconv.Atoi(src)

				if err == nil {
					wires[wire] = uint16(num)
				} else {
					if val, ok := wires[src]; ok {
						wires[wire] = val
					} else {
						wait = append(wait, line)
					}
				}
			}
		}

		if len(wait) == 0 {
			break
		} else {
			input = wait
		}
	}
	return wires
}

func part2(input []string, override uint16) map[string]uint16 {
	for i, line := range input {
		if strings.HasSuffix(line, "-> b") {
			input[i] = fmt.Sprintf("%d -> b", override)
		}
	}

	return part1(input)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	part1_ans := part1(lines)["a"]
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(lines, part1_ans)["a"]
	fmt.Printf("Part 2: %d\n", part2_ans)
}
