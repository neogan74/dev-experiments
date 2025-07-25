package maximumuniquesubarraysumafterdeletion

func maximumUniqueSubarraySumAfterDeletion(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxSum := 0
	currentSum := 0
	left := 0
	numCount := make(map[int]int)

	for right := 0; right < n; right++ {
		numCount[nums[right]]++
		currentSum += nums[right]

		for numCount[nums[right]] > 1 {
			numCount[nums[left]]--
			currentSum -= nums[left]
			left++
		}

		if left > 0 {
			currentSum -= nums[left-1]
			numCount[nums[left-1]]--
		}

		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}
