package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	// Start with a very high minimum price and 0 profit.
	minPrice := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		// Update the minimum price if the current price is lower.
		if price < minPrice {
			minPrice = price
		}

		// Calculate the current profit.
		profit := price - minPrice

		// Update the maximum profit if the current profit is higher.
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Maximum Profit:", maxProfit(prices1)) // Expected output: 5

	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println("Maximum Profit:", maxProfit(prices2)) // Expected output: 0
}
