package main

import (
	"bufio"
	"log"
	"os"
)

const (
	X = 'X'
	M = 'M'
	A = 'A'
	S = 'S'
)

type Point struct {
	X int
	Y int
}

type Matrix struct {
	matrix [][]byte
	xdim   int
	ydim   int
}

func NewMatrix(matrix [][]byte) *Matrix {
	return &Matrix{
		matrix: matrix,
		xdim:   len(matrix[0]),
		ydim:   len(matrix),
	}
}

func (m *Matrix) X() int {
	return m.xdim
}

func (m *Matrix) Y() int {
	return m.ydim
}

func (m *Matrix) Get(p Point) byte {
	if p.X < 0 || p.X >= m.xdim || p.Y < 0 || p.Y >= m.ydim {
		return 0
	}

	return m.matrix[p.Y][p.X]
}

func (m *Matrix) Xmas(pts []Point) int {
	var xmas = []byte{X, M, A, S}

	for i := 0; i < 4; i++ {
		if m.Get(pts[i]) != xmas[i] {
			return 0
		}
	}

	return 1
}

func (m *Matrix) Xmas2(mid Point) int {
	var mas1 = "MAS"
	var mas2 = "SAM"

	m1 := m.Get(mid)
	c1 := m.Get(Point{mid.X - 1, mid.Y - 1})
	c2 := m.Get(Point{mid.X + 1, mid.Y - 1})
	c3 := m.Get(Point{mid.X + 1, mid.Y + 1})
	c4 := m.Get(Point{mid.X - 1, mid.Y + 1})

	row1 := string([]byte{c1, m1, c3})
	row2 := string([]byte{c2, m1, c4})

	if (row1 == mas1 || row1 == mas2) && (row2 == mas1 || row2 == mas2) {
		return 1
	}

	return 0
}

func part1() {
	f, err := os.Open("cmd/day4/input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var data [][]byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, []byte(scanner.Text()))
	}

	mat := NewMatrix(data)
	cnt := 0

	for y := 0; y < mat.Y(); y++ {
		for x := 0; x < mat.X(); x++ {
			cnt += mat.Xmas([]Point{{x, y}, {x, y - 1}, {x, y - 2}, {x, y - 3}})             // north
			cnt += mat.Xmas([]Point{{x, y}, {x + 1, y - 1}, {x + 2, y - 2}, {x + 3, y - 3}}) // north-east
			cnt += mat.Xmas([]Point{{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y}})             // east
			cnt += mat.Xmas([]Point{{x, y}, {x + 1, y + 1}, {x + 2, y + 2}, {x + 3, y + 3}}) // south-east
			cnt += mat.Xmas([]Point{{x, y}, {x, y + 1}, {x, y + 2}, {x, y + 3}})             // south
			cnt += mat.Xmas([]Point{{x, y}, {x - 1, y + 1}, {x - 2, y + 2}, {x - 3, y + 3}}) // south-west
			cnt += mat.Xmas([]Point{{x, y}, {x - 1, y}, {x - 2, y}, {x - 3, y}})             // west
			cnt += mat.Xmas([]Point{{x, y}, {x - 1, y - 1}, {x - 2, y - 2}, {x - 3, y - 3}}) // north-west
		}
	}

	log.Printf("cnt: %d", cnt)
}

func main() {
	f, err := os.Open("cmd/day4/input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var data [][]byte

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data = append(data, []byte(scanner.Text()))
	}

	mat := NewMatrix(data)
	cnt := 0

	for y := 0; y < mat.Y(); y++ {
		for x := 0; x < mat.X(); x++ {
			cnt += mat.Xmas2(Point{x, y})
		}
	}

	log.Printf("cnt: %d", cnt)
}
