package _19_ContainsDuplicateII

// containsNearbyDuplicate проверяет, есть ли в nums две одинаковые цифры
// на расстоянии не более k.
func containsNearbyDuplicate(nums []int, k int) bool {
	last := make(map[int]int) // число → последний индекс
	for i, x := range nums {
		if prev, ok := last[x]; ok && i-prev <= k {
			return true
		}
		last[x] = i
	}
	return false
}
