package main

import (
	"fmt"
)

func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost, tank, start := 0, 0, 0, 0

	for i := 0; i < len(gas); i++ {
		totalGas += gas[i]
		totalCost += cost[i]

		tank += gas[i] - cost[i]

		// If the tank becomes negative, reset and move to next station
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}

	if totalGas < totalCost {
		return -1
	}

	return start
}

func main() {
	fmt.Println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) // Output: 3
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))             // Output: -1
}
