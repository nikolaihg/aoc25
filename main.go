package main

import (
	"fmt"

	"github.com/nikolaihg/aoc25/day02"
)

func main() {
	/*
		// day 1
		p1, p2 := day01.Solve("inputs/day01_test.txt")
		fmt.Printf("Part 1: %d\n", p1)
		fmt.Printf("Part 2: %d\n", p2)
	*/
	p1, p2 := day02.Solve("inputs/day02.txt")
	fmt.Printf("Part 1: %d\n", p1)
	fmt.Printf("Part 2: %d\n", p2)
}
