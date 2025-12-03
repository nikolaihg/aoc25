package utils

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
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

func ReadCSV(fp string) [][]string {
	f, err := os.Open(fp)
	if err != nil {
		log.Fatalf("Coud not open file: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	var records [][]string

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		records = append(records, line)
	}

	return records
}
