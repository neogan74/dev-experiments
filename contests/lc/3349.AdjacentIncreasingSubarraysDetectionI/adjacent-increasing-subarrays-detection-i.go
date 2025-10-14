package _349_AdjacentIncreasingSubarraysDetectionI

func hasUniqueElements(nums []int, k int) bool {
	n := len(nums)

	// Проверяем все возможные позиции для пары подмассивов
	// Первый подмассив: [i, i+k+1], второй [i+k, i+2*k - 1]
	for i := 0; i < n-2*k; i++ {
		if strictlyIncreaseing(nums[i:i+2*k], k) {
			return true
		}
	}
	return false
}

func strictlyIncreaseing(sub []int, k int) bool {
	return isInc(sub[:k]) && isInc(sub[k:])
}

func isInc(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
