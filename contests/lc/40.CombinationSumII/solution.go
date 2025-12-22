package combinationsumii

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates) // Step 1: Sort to handle duplicates

	var backtrack func(start int, currentTarget int, path []int)
	backtrack = func(start int, currentTarget int, path []int) {
		if currentTarget == 0 {
			// Found a valid combination; make a copy of path
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := start; i < len(candidates); i++ {
			// Step 2: Skip duplicate elements at the same level
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			// Step 3: Early termination if element exceeds remaining target
			if candidates[i] > currentTarget {
				break
			}

			path = append(path, candidates[i])
			// Step 4: Recursive call with next index i+1
			backtrack(i+1, currentTarget-candidates[i], path)
			// Step 5: Backtrack by removing the last element
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target, []int{})
	return res
}
