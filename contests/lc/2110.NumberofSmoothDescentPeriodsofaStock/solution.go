package numberofsmoothdescentperiodsofastock

// getDescentPeriods returns the number of smooth descent periods.
// A smooth descent run of length k contributes k*(k+1)/2 periods.
func getDescentPeriods(prices []int) int64 {
	var total int64
	runLength := int64(0)

	for i := 0; i < len(prices); i++ {
		if i > 0 && prices[i-1]-prices[i] == 1 {
			runLength++
		} else {
			// End of previous run.
			total += runLength * (runLength + 1) / 2
			runLength = 1
		}
	}

	// Account for the final run.
	total += runLength * (runLength + 1) / 2
	return total
}
