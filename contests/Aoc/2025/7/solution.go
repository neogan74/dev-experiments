package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

// countSplits simulates the downward beam and returns how many times it hits a splitter.
func countSplits(grid []string) int {
	if len(grid) == 0 {
		return 0
	}

	width := len(grid[0])
	startRow, startCol, found := 0, 0, false
	for r, row := range grid {
		if c := strings.IndexByte(row, 'S'); c != -1 {
			startRow, startCol, found = r, c, true
			break
		}
	}
	if !found {
		return 0
	}

	beams := map[int]struct{}{startCol: {}}
	splits := 0

	for r := startRow + 1; r < len(grid); r++ {
		next := make(map[int]struct{})
		row := grid[r]

		for c := range beams {
			cell := row[c]
			if cell == '^' {
				splits++
				if c > 0 {
					next[c-1] = struct{}{}
				}
				if c+1 < width {
					next[c+1] = struct{}{}
				}
				continue
			}
			next[c] = struct{}{}
		}

		beams = next
		if len(beams) == 0 {
			break // all beams have exited the grid
		}
	}

	return splits
}

// countTimelines returns how many distinct timelines exist after the particle traverses the grid.
func countTimelines(grid []string) *big.Int {
	if len(grid) == 0 {
		return big.NewInt(0)
	}

	width := len(grid[0])
	startRow, startCol, found := 0, 0, false
	for r, row := range grid {
		if c := strings.IndexByte(row, 'S'); c != -1 {
			startRow, startCol, found = r, c, true
			break
		}
	}
	if !found {
		return big.NewInt(0)
	}

	beams := map[int]*big.Int{startCol: big.NewInt(1)}

	for r := startRow + 1; r < len(grid); r++ {
		next := make(map[int]*big.Int)
		row := grid[r]

		for c, count := range beams {
			cell := row[c]
			if cell == '^' {
				if c > 0 {
					addCount(next, c-1, count)
				}
				if c+1 < width {
					addCount(next, c+1, count)
				}
				continue
			}
			addCount(next, c, count)
		}

		beams = next
		if len(beams) == 0 {
			break
		}
	}

	total := big.NewInt(0)
	for _, count := range beams {
		total.Add(total, count)
	}
	return total
}

func addCount(m map[int]*big.Int, col int, add *big.Int) {
	if col < 0 {
		return
	}
	if existing := m[col]; existing != nil {
		existing.Add(existing, add)
		return
	}
	m[col] = new(big.Int).Set(add)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func main() {
	inputPath := "input.txt"
	flag.StringVar(&inputPath, "input", inputPath, "path to input file")
	part2 := flag.Bool("part2", false, "solve part 2 (timeline count)")
	flag.Parse()
	if flag.NArg() > 0 && flag.Arg(0) != "" {
		inputPath = flag.Arg(0) // allow positional override for compatibility
	}

	grid, err := readLines(inputPath)
	if err != nil {
		log.Fatalf("read input %q: %v", inputPath, err)
	}

	if *part2 {
		fmt.Println(countTimelines(grid))
		return
	}

	fmt.Println(countSplits(grid))
}
