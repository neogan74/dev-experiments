package maximizesubarraysafterremovingoneconflictingpair

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	valid := int64(0)
	maxLeft, second := 0, 0
	gains := make([]int64, n+2)
	conflicts := make([][]int, n+2)
	for _, p := range conflictingPairs {
		a, b := p[0], p[1]
		if a > b {
			a, b = b, a
		}
		conflicts[b] = append(conflicts[b], a)
	}
	for right := 1; right <= n; right++ {
		for _, left := range conflicts[right] {
			if left > maxLeft {
				second = maxLeft
				maxLeft = left
			} else if left > second {
				second = left
			}
		}
		valid += int64(right - maxLeft)
		gains[maxLeft] += int64(maxLeft - second)
	}
	maxGain := int64(0)
	for _, g := range gains {
		if g > maxGain {
			maxGain = g
		}
	}
	return valid + maxGain
}
