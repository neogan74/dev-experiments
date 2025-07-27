package counthillsandvalleysinanarray

func countHillValley(nums []int) int {
	// Убираем повторы подряд
	clean := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			clean = append(clean, nums[i])
		}
	}

	count := 0
	for i := 1; i < len(clean)-1; i++ {
		if clean[i] > clean[i-1] && clean[i] > clean[i+1] {
			count++ // hill
		} else if clean[i] < clean[i-1] && clean[i] < clean[i+1] {
			count++ // valley
		}
	}

	return count
}
