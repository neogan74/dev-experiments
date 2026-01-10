package main

func minimumDeleteSum(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)

	dp := make([]int, n+1)
	for j := n - 1; j >= 0; j-- {
		dp[j] = dp[j+1] + int(s2[j])
	}

	for i := m - 1; i >= 0; i-- {
		prevDiag := dp[n]
		dp[n] = dp[n] + int(s1[i])
		for j := n - 1; j >= 0; j-- {
			temp := dp[j]
			if s1[i] == s2[j] {
				dp[j] = prevDiag
			} else {
				del1 := int(s1[i]) + temp
				del2 := int(s2[j]) + dp[j+1]
				if del1 < del2 {
					dp[j] = del1
				} else {
					dp[j] = del2
				}
			}
			prevDiag = temp
		}
	}

	return dp[0]
}
