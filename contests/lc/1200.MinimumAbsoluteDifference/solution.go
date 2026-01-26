package minimumabsolutedifference

import "sort"

func minimumAbsDifference(arr []int) [][]int {

	sort.Ints(arr)

	minDiff := int(1e9)
	result := [][]int{}

	// First pass
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	// Second pass
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minDiff {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}

	return result
}
