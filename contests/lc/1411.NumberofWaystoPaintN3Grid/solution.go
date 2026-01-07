package numberofwaystopaintn3grid

func numOfWays(n int) int {
	const mod int64 = 1000000007
	a, b := int64(6), int64(6)
	for i := 2; i <= n; i++ {
		na := (a*3 + b*2) % mod
		nb := (a*2 + b*2) % mod
		a, b = na, nb
	}
	return int((a + b) % mod)
}
