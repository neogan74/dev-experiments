package _28_LongestConsecutiveSequence

// longestConsecutive возвращает длину самой длинной последовательности подряд идущих чисел.
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := make(map[int]bool, len(nums))
	for _, x := range nums {
		set[x] = true
	}

	maxLen := 0
	for x := range set {
		// если нет x-1, значит начинаем новую последовательность
		if !set[x-1] {
			length := 1
			for y := x + 1; set[y]; y++ {
				length++
			}
			if length > maxLen {
				maxLen = length
			}
		}
	}
	return maxLen
}
