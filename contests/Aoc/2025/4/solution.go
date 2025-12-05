package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const defaultInput = "input.txt"

func main() {
	var part2 bool

	root := &cobra.Command{
		Use:   "day4 [input]",
		Short: "Compute accessible paper rolls (part1) or total removable (part2)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path := defaultInput
			if len(args) == 1 {
				path = args[0]
			}

			grid, err := readGrid(path)
			if err != nil {
				return err
			}

			if part2 {
				fmt.Println(countRemovable(grid))
				return nil
			}

			fmt.Println(countAccessible(grid))
			return nil
		},
	}

	root.Flags().BoolVar(&part2, "part2", false, "run part two calculation")

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func readGrid(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var grid []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func countAccessible(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for r := 0; r < rows; r++ {
		cols := len(grid[r])
		for c := 0; c < cols; c++ {
			if grid[r][c] != '@' {
				continue
			}
			adj := 0
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= rows {
					continue
				}
				if nc < 0 || nc >= len(grid[nr]) {
					continue
				}
				if grid[nr][nc] == '@' {
					adj++
				}
			}
			if adj < 4 {
				count++
			}
		}
	}
	return count
}

func countRemovable(grid []string) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}

	// Work on a mutable copy.
	current := make([][]byte, rows)
	for i, row := range grid {
		current[i] = []byte(row)
	}

	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	totalRemoved := 0
	for {
		var toRemove [][2]int
		for r := 0; r < rows; r++ {
			for c := 0; c < len(current[r]); c++ {
				if current[r][c] != '@' {
					continue
				}
				adj := 0
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr < 0 || nr >= rows {
						continue
					}
					if nc < 0 || nc >= len(current[nr]) {
						continue
					}
					if current[nr][nc] == '@' {
						adj++
					}
				}
				if adj < 4 {
					toRemove = append(toRemove, [2]int{r, c})
				}
			}
		}

		if len(toRemove) == 0 {
			break
		}

		totalRemoved += len(toRemove)
		for _, rc := range toRemove {
			current[rc[0]][rc[1]] = '.' // remove roll
		}
	}

	return totalRemoved
}
