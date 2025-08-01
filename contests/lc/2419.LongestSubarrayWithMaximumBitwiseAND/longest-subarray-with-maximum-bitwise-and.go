package longestsubarraywithmaximumbitwiseand

import "math"

// LongestSubarrayWithMaximumBitwiseAnd finds the length of the longest subarray
// where the bitwise AND of all elements is equal to the maximum element in that subarray
func LongestSubarrayWithMaximumBitwiseAnd(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxElement := math.MinInt32
	for _, num := range arr {
		if num > maxElement {
			maxElement = num
		}
	}

	maxLength := 0
	currentLength := 0

	for _, num := range arr {
		if num == maxElement {
			currentLength++
			if currentLength > maxLength {
				maxLength = currentLength
			}
		} else {
			currentLength = 0
		}
	}

	return maxLength
}
