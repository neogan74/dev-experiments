package rangeproductqueriesofpowers

func productQueries(n int, queries [][]int) []int {
	const mod = 1_000_000_007
	powers := []int{}
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			powers = append(powers, 1<<i)
		}
	}

	results := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		product := 1
		for j := left; j <= right; j++ {
			product = (product * powers[j]) % mod
		}
		results[i] = product
	}

	return results
}
