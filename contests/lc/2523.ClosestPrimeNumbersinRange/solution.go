package closestprimenumbersinrange

func closestPrimes(left int, right int) []int {
	if right < 2 {
		return []int{-1, -1}
	}

	isComposite := make([]bool, right+1)
	for i := 2; i*i <= right; i++ {
		if !isComposite[i] {
			for j := i * i; j <= right; j += i {
				isComposite[j] = true
			}
		}
	}

	prev := -1
	bestGap := int(1e9)
	ans := []int{-1, -1}

	for i := max(left, 2); i <= right; i++ {
		if !isComposite[i] {
			if prev != -1 {
				gap := i - prev
				if gap < bestGap {
					bestGap = gap
					ans[0], ans[1] = prev, i
				}
			}
			prev = i
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
