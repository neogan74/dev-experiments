package main

import "fmt"

func minimumIndex(nums []int) int {
	n := len(nums)
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	leftFreq := make(map[int]int)

	for i, num := range nums {
		leftFreq[num]++

		totalCount := freq[num]
		leftCount := leftFreq[num]
		rightCount := totalCount - leftCount

		if leftCount > (i+1)/2 && rightCount > (n-i-1)/2 {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{1, 2, 2, 2}
	res := minimumIndex(nums) // 2
	fmt.Println(res)

	nums2 := []int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}
	res = minimumIndex(nums2) // expected 4
	fmt.Println(res)

}
