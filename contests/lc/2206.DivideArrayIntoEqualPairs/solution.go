package main

import "fmt"

// divideArray checks if the array can be divided into pairs such that each pair consists of equal elements.
func divideArray(nums []int) bool {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}

	// Check that each frequency is even.
	for _, count := range freq {
		if count%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	// Example 1:
	nums1 := []int{3, 2, 3, 2, 2, 2}
	fmt.Println(divideArray(nums1)) // Expected output: true

	// Example 2:
	nums2 := []int{1, 2, 3, 4}
	fmt.Println(divideArray(nums2)) // Expected output: false
}
