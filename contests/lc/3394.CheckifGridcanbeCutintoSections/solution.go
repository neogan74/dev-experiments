package main

import (
	"fmt"
	"sort"
)

func canCutGrid(grid []string, horizontalCuts int, verticalCuts int) bool {
	rows, cols := len(grid), len(grid[0])
	prefix := make([][]int, rows+1)
	for i := range prefix {
		prefix[i] = make([]int, cols+1)
	}

	// Build 2D prefix sum
	total := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			val := int(grid[r][c] - '0')
			total += val
			prefix[r+1][c+1] = val + prefix[r][c+1] + prefix[r+1][c] - prefix[r][c]
		}
	}

	h, v := horizontalCuts+1, verticalCuts+1
	if total%(h*v) != 0 {
		return false
	}
	onesPerBlock := total / (h * v)
	onesPerRowSlice := total / h
	onesPerColSlice := total / v

	// Find row cuts
	rowCuts := []int{}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			count += int(grid[r][c] - '0')
		}
		if count == onesPerRowSlice {
			rowCuts = append(rowCuts, r)
			count = 0
		} else if count > onesPerRowSlice {
			return false
		}
	}
	if len(rowCuts) != h {
		return false
	}

	// Find column cuts
	colCuts := []int{}
	count = 0
	for c := 0; c < cols; c++ {
		for r := 0; r < rows; r++ {
			count += int(grid[r][c] - '0')
		}
		if count == onesPerColSlice {
			colCuts = append(colCuts, c)
			count = 0
		} else if count > onesPerColSlice {
			return false
		}
	}
	if len(colCuts) != v {
		return false
	}

	// Check each block
	rStart := 0
	for i := 0; i < h; i++ {
		rEnd := rowCuts[i]
		cStart := 0
		for j := 0; j < v; j++ {
			cEnd := colCuts[j]
			ones := getSum(prefix, rStart, cStart, rEnd, cEnd)
			if ones != onesPerBlock {
				return false
			}
			cStart = cEnd + 1
		}
		rStart = rEnd + 1
	}

	return true
}

// getSum
func getSum(prefix [][]int, r1, c1, r2, c2 int) int {
	return prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
}

// checkValidCuts
func checkValidCuts(n int, rectangles [][]int) bool {
	var xIntervals, yIntervals [][]int

	for _, rect := range rectangles {
		xIntervals = append(xIntervals, []int{rect[0], rect[2]})
		yIntervals = append(yIntervals, []int{rect[1], rect[3]})
	}

	return check(xIntervals) || check(yIntervals)
}

func check(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	sections := 0
	maxEnd := intervals[0][1]

	for _, interval := range intervals {
		if maxEnd <= interval[0] {
			sections++
		}
		maxEnd = max(maxEnd, interval[1])
	}

	return sections >= 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	a := canCutGrid([]string{"101101", "111111", "000111", "110100", "110000", "110110", "110100", "011000", "000000", "011111", "000110", "001000", "110101", "101000", "100010"}, 5, 5)
	fmt.Println(a)
	checkValidCuts(10, [][]int{{0, 0, 4, 4}, {0, 0, 1, 1}, {1, 0, 2, 1}, {1, 2, 3, 0}, {2, 1, 3, 2}, {2, 2, 4, 3}, {2, 3, 4, 4}})
}
