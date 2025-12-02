package day01

import (
	"strconv"

	"github.com/nikolaihg/aoc25/utils"
)

func Solve(path string) (int, int) {
	input := utils.ReadLines(path)
	p1 := solveFirst(input)
	p2 := solveSecond(input)
	return p1, p2
}

/* first version of part 1
func solveFirst(input []string) int {
	pos := 50
	var code int
	for _, line := range input {
		if strings.HasPrefix(line, "R") {
			c, _ := strconv.Atoi(line[1:])
			pos = pos + c
		} else if strings.HasPrefix(line, "L") {
			c, _ := strconv.Atoi(line[1:])
			pos = pos - c
		}
		if pos%100 == 0 {
			code++
		}
	}
	return code
}
*/

// improved version after solving part 2
func solveFirst(input []string) int {
	pos := 50
	var code int
	delta := map[byte]int{'L': -1, 'R': 1}
	for _, line := range input {
		c, _ := strconv.Atoi(line[1:])
		// multiple pos & (1/-1) clicks
		pos += delta[line[0]] * c
		if pos%100 == 0 {
			code++
		}

	}
	return code
}

func solveSecond(input []string) int {
	pos := 50
	var code int
	delta := map[byte]int{'L': -1, 'R': 1}
	for _, line := range input {
		c, _ := strconv.Atoi(line[1:])
		// check if current position + [1..line] == 0
		for range c {
			if pos += delta[line[0]]; pos%100 == 0 {
				code++
			}
		}
	}
	return code
}
