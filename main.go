package main

import (
	"fmt"

	"github.com/nikolaihg/aoc25/day01"
)

func main() {
	fp := "inputs/day01.txt"
	p1, p2 := day01.Solve(fp)
	fmt.Println("day01 part1:", p1)
	fmt.Println("day01 part2:", p2)
}
