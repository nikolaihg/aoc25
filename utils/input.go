package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadLines(fp string) []string {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatalf("Coud not open file: %v", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("Error during scanning: %v", err)
	}
	return lines
}

func SplitString(s string) []string {
	lines := strings.Split(s, "\n")
	var out []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		out = append(out, line)
	}
	return out
}

func SplitByBlank(s string) [][]string {
	lines := strings.Split(s, "\n")
	var groups [][]string
	var cur []string

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			if len(cur) > 0 {
				groups = append(groups, cur)
				cur = nil
			}
			continue
		}
		cur = append(cur, l)
	}
	if len(cur) > 0 {
		groups = append(groups, cur)
	}
	return groups
}

func Mod(a int, b int) int {
	return (a%b + b) % b
}
