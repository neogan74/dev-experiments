package _55_DeleteColumnstoMakeSortedII

func minDeletionSize(strs []string) int {
	n := len(strs)
	m := len(strs[0])

	fixed := make([]bool, n-1)
	deletions := 0

	allFixed := func() bool {
		for _, v := range fixed {
			if !v {
				return false
			}
		}
		return true
	}

	for col := 0; col < m; col++ {
		bad := false
		for i := 0; i < n-1; i++ {
			if !fixed[i] && strs[i][col] > strs[i+1][col] {
				bad = true
				break
			}
		}
		if bad {
			deletions++
			continue
		}

		for i := 0; i < n-1; i++ {
			if !fixed[i] && strs[i][col] < strs[i+1][col] {
				fixed[i] = true
			}
		}
		if allFixed() {
			break
		}

	}
	return deletions
}
