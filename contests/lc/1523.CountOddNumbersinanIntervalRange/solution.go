package countoddnumbersinanintervalrange

func countOdds(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}

func countOdds2(low int, high int) int {
	if high%2 == 0 && low%2 == 0 {
		return (high - low) / 2
	}
	return ((high - low) / 2) + 1
}
