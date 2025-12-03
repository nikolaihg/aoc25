package day02

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(path string) (int, int) {
	input, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	data := strings.Split(strings.TrimSpace(string(input)), ",")

	p1 := solveFirst(data)
	p2 := solveSecond(data)
	return p1, p2
}

func solveFirst(data []string) int {
	sum := 0
	for _, s := range data {
		parts := strings.Split(s, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		for i := a; i <= b; i++ {
			numStr := strconv.Itoa(i)
			if isRepeated(numStr) {
				sum += i
			}
		}
	}
	return sum
}

func solveSecond(data []string) int {
	sum := 0
	for _, s := range data {
		parts := strings.Split(s, "-")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		for i := a; i <= b; i++ {
			numStr := strconv.Itoa(i)
			if twiceOrMore(numStr) {
				sum += i
			}
		}
	}
	return sum
}

func isRepeated(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]

	return firstHalf == secondHalf
}

// isRepeated twice or more
func twiceOrMore(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}
		pattern := s[:i]
		repetitions := n / i

		if repetitions >= 2 {
			if s == strings.Repeat(pattern, repetitions) {
				return true
			}
		}
	}
	return false
}
