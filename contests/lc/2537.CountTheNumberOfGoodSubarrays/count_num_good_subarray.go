package main

func countGood(nums []int, k int) int64 {
	freq := make(map[int]int)
	var ans int64
	var cur int
	left := 0

	for right, x := range nums {
		cur += freq[x]
		freq[x]++

		for cur >= k {
			ans += int64(len(nums) - right)
			freq[nums[left]]--
			cur -= freq[nums[left]]
			left++
		}
	}

	return ans
}
