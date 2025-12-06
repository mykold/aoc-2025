package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func iat(grid []rune, i int, eol int) ([8]int, bool) {
	n := [8]int{
		-1, -1, -1,
		-1 /* */, -1,
		-1, -1, -1,
	}

	l := i%eol != 0
	r := (i+1)%eol != 0

	// n[0] n[1] n[2]
	// n[3]   i  n[4]
	// n[5] n[6] n[7]

	{
		j := i - eol - 1
		if l && j >= 0 {
			n[0] = j
		}
	}

	{
		j := i - eol
		if j >= 0 {
			n[1] = j
		}
	}

	{
		j := i - eol + 1
		if r && j >= 0 {
			n[2] = j
		}
	}

	{
		j := i - 1
		if l && j >= 0 {
			n[3] = j
		}
	}

	// i

	{
		j := i + 1
		if r && j < len(grid) {
			n[4] = j
		}
	}

	{
		j := i + eol - 1
		if l && j < len(grid) {
			n[5] = j
		}
	}

	{
		j := i + eol
		if j < len(grid) {
			n[6] = j
		}
	}

	{
		j := i + eol + 1
		if r && j < len(grid) {
			n[7] = j
		}
	}

	return n, true
}

func at(grid []rune, i int, eol int) ([8]rune, bool) {
	n := [8]rune{
		-1, -1, -1,
		-1 /* */, -1,
		-1, -1, -1,
	}

	ps, ok := iat(grid, i, eol)
	if !ok {
		return n, false
	}

	for j, k := range ps {
		if k == -1 {
			continue
		}
		n[j] = grid[k]
	}

	return n, true
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	eol := 0
	grid := make([]rune, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if eol == 0 {
			eol = len(line)
		}

		grid = append(grid, []rune(line)...)
	}

	var r int

	for i := range len(grid) {
		// Count only for '@' positions
		if grid[i] != '@' {
			continue
		}

		// Get the rolls of paper in the eight adjacent positions
		ps, ok := at(grid, i, eol)
		if !ok {
			log.Fatal(ps)
		}

		// The count of rolls of paper in the eight adjacent positions
		var c int
		for _, p := range ps {
			if p == '@' {
				c++
			}
		}

		// The forklifts can only access a roll of paper if there are fewer
		// than four rolls of paper in the eight adjacent positions.
		if c < 4 {
			r += 1
		}
	}

	fmt.Println(r)
}
