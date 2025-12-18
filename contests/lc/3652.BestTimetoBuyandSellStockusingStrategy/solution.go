package besttimetobuyanssellstockusingstrategy

// maxProfit returns the maximum profit after at most one window modification.
func maxProfit(prices []int, strategy []int, k int) int64 {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	half := k / 2

	prefixPrice := make([]int64, n+1)
	prefixProd := make([]int64, n+1) // strategy[i] * prices[i]

	for i := 0; i < n; i++ {
		p := int64(prices[i])
		prefixPrice[i+1] = prefixPrice[i] + p
		prefixProd[i+1] = prefixProd[i] + p*int64(strategy[i])
	}

	base := prefixProd[n]
	var bestDelta int64

	for start := 0; start+k <= n; start++ {
		mid := start + half
		end := start + k

		firstHalf := prefixProd[mid] - prefixProd[start]         // strategy * price
		secondHalfPrice := prefixPrice[end] - prefixPrice[mid]   // price
		secondHalfProd := prefixProd[end] - prefixProd[mid]      // strategy * price
		delta := -firstHalf + (secondHalfPrice - secondHalfProd) // gain if window modified

		if delta > bestDelta {
			bestDelta = delta
		}
	}

	return int64(base + bestDelta)
}
