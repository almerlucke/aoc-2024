package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func countOccurence(value int64, list []int64) int64 {
	var cnt int64
	for _, v := range list {
		if v == value {
			cnt++
		}
	}
	return cnt
}

func main() {
	f, err := os.Open("cmd/day1/input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var leftNumbers []int64
	var rightNumbers []int64

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")

		leftNum, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		rightNum, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		leftNumbers = append(leftNumbers, leftNum)
		rightNumbers = append(rightNumbers, rightNum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(leftNumbers)
	slices.Sort(rightNumbers)

	var totalDistance int64
	var similarityScore int64

	for i, left := range leftNumbers {
		distance := rightNumbers[i] - left
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
		count := countOccurence(left, rightNumbers)
		similarityScore += count * left
	}

	log.Printf("total distance = %d", totalDistance)
	log.Printf("similarityScore = %d", similarityScore)
}
