package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	profit := 0
	// Start iterating from the second day
	for i := 1; i < len(prices); i++ {
		// If today's price is higher than yesterday's, capture the profit.
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Maximum Profit:", maxProfit(prices1)) // Expected output: 7

	prices2 := []int{1, 2, 3, 4, 5}
	fmt.Println("Maximum Profit:", maxProfit(prices2)) // Expected output: 4

	prices3 := []int{7, 6, 4, 3, 1}
	fmt.Println("Maximum Profit:", maxProfit(prices3)) // Expected output: 0
}
