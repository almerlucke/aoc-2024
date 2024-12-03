package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func distance(x int, y int) int {
	d := x - y
	if d < 0 {
		return -d
	}
	return d
}

func recordMinusOne(record []int, index int) []int {
	var minus []int
	for i, v := range record {
		if i == index {
			continue
		}
		minus = append(minus, v)
	}
	return minus
}

func isValidRecord(record []int) bool {
	if isValidRecordIntern(record) {
		return true
	}

	for i := range len(record) {
		minus := recordMinusOne(record, i)
		if isValidRecordIntern(minus) {
			return true
		}
	}

	return false
}

func isValidRecordIntern(record []int) bool {
	var direction int
	var prevLevel int

	for i, level := range record {
		if i > 0 {
			d := distance(level, prevLevel)
			if d < 1 || d > 3 {
				return false
			}

			if i == 1 {
				if level < prevLevel {
					direction = -1
				} else {
					direction = 1
				}
			} else if level < prevLevel {
				if direction == 1 {
					return false
				}
			} else if direction == -1 {
				return false
			}
		}

		prevLevel = level
	}

	return true
}

func main() {
	f, err := os.Open("cmd/day2/input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var numValidRecords int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		record := make([]int, len(parts))
		for i, p := range parts {
			num, _ := strconv.ParseInt(p, 10, 64)
			record[i] = int(num)
		}
		valid := isValidRecord(record)
		log.Printf("record: %v -> is valid = %t", record, valid)
		if valid {
			numValidRecords++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Valid records: %d", numValidRecords)
}
