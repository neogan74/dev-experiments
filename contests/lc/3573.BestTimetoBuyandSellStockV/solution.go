package _besttimeofbiyandsellstockv

// maximumProfit computes the maximum profit with at most k transactions where
// each transaction can be a long (buy then sell) or short (sell then buy back)
// position. We track profits in three states for each completed transaction count.
func maximumProfit(prices []int, k int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}
	if limit := len(prices) / 2; k > limit {
		k = limit
	}

	const neg int64 = -1 << 60
	rest := make([]int64, k+1)  // no open position
	long := make([]int64, k+1)  // currently long
	short := make([]int64, k+1) // currently short
	for i := 1; i <= k; i++ {
		rest[i] = neg
	}
	for i := 0; i <= k; i++ {
		long[i], short[i] = neg, neg
	}
	rest[0] = 0

	for _, price := range prices {
		p := int64(price)
		nextRest := make([]int64, k+1)
		nextLong := make([]int64, k+1)
		nextShort := make([]int64, k+1)
		copy(nextRest, rest)
		copy(nextLong, long)
		copy(nextShort, short)

		for t := 0; t <= k; t++ {
			if long[t] != neg && t+1 <= k {
				if val := long[t] + p; val > nextRest[t+1] {
					nextRest[t+1] = val
				}
			}
			if short[t] != neg && t+1 <= k {
				if val := short[t] - p; val > nextRest[t+1] {
					nextRest[t+1] = val
				}
			}
			if rest[t] != neg {
				if val := rest[t] - p; val > nextLong[t] {
					nextLong[t] = val
				}
				if val := rest[t] + p; val > nextShort[t] {
					nextShort[t] = val
				}
			}
		}

		rest, long, short = nextRest, nextLong, nextShort
	}

	ans := rest[0]
	for _, v := range rest {
		if v > ans {
			ans = v
		}
	}
	return int(ans)
}
