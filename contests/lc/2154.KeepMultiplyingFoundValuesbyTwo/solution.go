package keepmultiplyingfoundvaluesbytwo

func findFinalValue(nums []int, original int) int {
	numsMap := make(map[int]bool)
	for _, num := range nums {
		numsMap[num] = true
	}

	for {
		if _, exists := numsMap[original]; exists {
			original *= 2
		} else {
			break
		}
	}

	return original
}
