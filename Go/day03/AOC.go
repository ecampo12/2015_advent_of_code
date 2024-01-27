package main

import (
	"fmt"
	"os"
)

type coord struct {
	x, y int
}

type Santas struct {
	coord
	visted map[coord]bool
}

func part1(input []byte) int {
	santa := Santas{coord{0, 0}, map[coord]bool{{0, 0}: true}}
	for _, c := range input {
		switch c {
		case '^':
			santa.y++
		case 'v':
			santa.y--
		case '>':
			santa.x++
		case '<':
			santa.x--
		}

		santa.visted[santa.coord] = true
	}
	return len(santa.visted)
}

func intersection(s1 Santas, s2 Santas) int {
	intersect := 0
	for k := range s1.visted {
		if _, ok := s2.visted[k]; ok {
			intersect++
		}
	}
	return intersect
}

func part2(input []byte) int {
	s := []Santas{
		{coord{0, 0}, map[coord]bool{{0, 0}: true}}, // santa
		{coord{0, 0}, map[coord]bool{{0, 0}: true}}, // robo-santa
	}

	for i, c := range input {
		var d *Santas
		if i%2 == 0 {
			d = &s[0]
		} else {
			d = &s[1]
		}

		switch c {
		case '^':
			d.y++
		case 'v':
			d.y--
		case '>':
			d.x++
		case '<':
			d.x--
		}
		d.visted[d.coord] = true
	}
	return len(s[0].visted) + len(s[1].visted) - intersection(s[0], s[1])
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print("Error reading file\n")
		panic(err)
	}

	part1_ans := part1(input)
	fmt.Printf("Part 1: %d\n", part1_ans)

	part2_ans := part2(input)
	fmt.Printf("Part 2: %d\n", part2_ans)
}
