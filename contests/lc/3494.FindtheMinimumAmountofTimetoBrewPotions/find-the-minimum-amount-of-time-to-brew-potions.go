package _494_FindtheMinimumAmountofTimetoBrewPotions

func minTime(skill []int, mana []int) int64 {
	n, m := len(skill), len(mana)
	if n == 0 || m == 0 {
		return 0
	}
	dp := make([]int64, n)
	for _, manaVal := range mana {
		prevMachineDone := int64(0)
		for i, s := range skill {
			timeOnMachine := int64(s) * int64(manaVal)
			if dp[i] > prevMachineDone {
				prevMachineDone = dp[i]
			}
			dp[i] = prevMachineDone + timeOnMachine
			prevMachineDone = dp[i]
		}
	}
	return dp[n-1]
}
