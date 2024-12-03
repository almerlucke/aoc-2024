package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func part1() {
	f, err := os.Open("cmd/day3/input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	reg := regexp.MustCompile(`mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`)
	var total int64
	for _, match := range reg.FindAllStringSubmatch(string(data), -1) {
		n1, _ := strconv.ParseInt(match[1], 10, 64)
		n2, _ := strconv.ParseInt(match[2], 10, 64)
		total += n1 * n2
	}

	log.Printf("total: %d", total)
}

type Entity interface {
	Index() int
}

type MulMatch struct {
	index int
	mult  int64
}

func (m *MulMatch) Index() int {
	return m.index
}

type Do struct {
	index int
	Do    bool
}

func (d *Do) Index() int {
	return d.index
}

func main() {
	f, err := os.Open("cmd/day3/input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var entities []Entity

	reg := regexp.MustCompile(`mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`)
	for _, loc := range reg.FindAllSubmatchIndex(data, -1) {
		n1, _ := strconv.ParseInt(string(data[loc[2]:loc[3]]), 10, 64)
		n2, _ := strconv.ParseInt(string(data[loc[4]:loc[5]]), 10, 64)

		entities = append(entities, &MulMatch{
			index: loc[0],
			mult:  n1 * n2,
		})
	}

	reg = regexp.MustCompile(`don't\(\)`)
	for _, loc := range reg.FindAllSubmatchIndex(data, -1) {
		entities = append(entities, &Do{
			index: loc[0],
			Do:    false,
		})
	}

	reg = regexp.MustCompile(`do\(\)`)
	for _, loc := range reg.FindAllSubmatchIndex(data, -1) {
		entities = append(entities, &Do{
			index: loc[0],
			Do:    true,
		})
	}

	slices.SortFunc(entities, func(i, j Entity) int {
		return i.Index() - j.Index()
	})

	var (
		total   int64
		enabled = true
	)

	for _, entity := range entities {
		switch t := entity.(type) {
		case *MulMatch:
			if enabled {
				total += t.mult
			}
		case *Do:
			enabled = t.Do
		}
	}

	log.Printf("total: %d", total)
}
