package repeatedElementinSize2NArray

func repeatedNTimes(nums []int) int {
	seen := make(map[int]bool, len(nums))
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return 0
}
