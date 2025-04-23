package main

import (
	"fmt"
	"math"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	values := []int{}

	for _, row := range grid {
		values = append(values, row...)
	}

	sort.Ints(values)

	for _, val := range values {
		if int(math.Abs(float64(val-values[0])))%x != 0 {
			return -1
		}
	}

	median := values[len(values)/2]
	operations := 0

	for _, val := range values {
		operations += int(math.Abs(float64(val-median))) / x
	}

	return operations
}

func main() {
	grid := [][]int{{2, 4}, {6, 8}}
	x := 2
	res := minOperations(grid, x)
	fmt.Println(res)
}
