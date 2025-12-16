package maximumprofitfromtradingstockswithdiscounts

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	tree := make([][]int, n)
	for _, e := range hierarchy {
		tree[e[0]-1] = append(tree[e[0]-1], e[1]-1)
	}

	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, budget+1)
		}
	}

	merge := func(A, B []int) []int {
		C := make([]int, budget+1)
		for i := range C {
			C[i] = -1e9
		}
		for i := 0; i <= budget; i++ {
			if A[i] < 0 {
				continue
			}
			for j := 0; i+j <= budget; j++ {
				if C[i+j] < A[i]+B[j] {
					C[i+j] = A[i] + B[j]
				}
			}
		}
		return C
	}

	var dfs func(int)
	dfs = func(u int) {
		for _, v := range tree[u] {
			dfs(v)
		}

		for pb := 0; pb <= 1; pb++ {
			price := present[u]
			if pb == 1 {
				price /= 2
			}
			profit := future[u] - price

			skip := make([]int, budget+1)
			for _, v := range tree[u] {
				skip = merge(skip, dp[v][0])
			}

			take := make([]int, budget+1)
			for i := range take {
				take[i] = -1e9
			}

			if price <= budget {
				base := make([]int, budget+1)
				for _, v := range tree[u] {
					base = merge(base, dp[v][1])
				}
				for b := price; b <= budget; b++ {
					take[b] = base[b-price] + profit
				}
			}

			for b := 0; b <= budget; b++ {
				dp[u][pb][b] = max(skip[b], take[b])
			}
		}
	}

	dfs(0)
	ans := 0
	for _, v := range dp[0][0] {
		if v > ans {
			ans = v
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

func maxProfit2(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	// build tree (0-indexed)
	g := make([][]int, n)
	for _, e := range hierarchy {
		u, v := e[0]-1, e[1]-1
		g[u] = append(g[u], v)
	}

	const NEG = -1 << 60

	// dfs returns (dp0, dp1)
	// dp0[c] = max profit in subtree(u) with total cost c when parent of u DID NOT buy
	// dp1[c] = max profit in subtree(u) with total cost c when parent of u DID buy
	var dfs func(u int) ([]int, []int)
	dfs = func(u int) ([]int, []int) {
		// base arrays
		// start with base = children combined assuming u DOES NOT buy (children see parent NOT bought)
		base := make([]int, budget+1)
		for i := 0; i <= budget; i++ {
			base[i] = NEG
		}
		base[0] = 0

		// collect children dp (we need both child.dp0 and child.dp1)
		children := g[u]
		childDP0 := make([][]int, 0, len(children))
		childDP1 := make([][]int, 0, len(children))

		for _, v := range children {
			c0, c1 := dfs(v)
			childDP0 = append(childDP0, c0)
			childDP1 = append(childDP1, c1)

			// merge base with child's dp0  (since u currently NOT buying -> child sees parent NOT bought)
			newBase := make([]int, budget+1)
			for i := 0; i <= budget; i++ {
				newBase[i] = NEG
			}
			for costU := 0; costU <= budget; costU++ {
				if base[costU] == NEG {
					continue
				}
				valU := base[costU]
				// try all costV from child.dp0
				for costV := 0; costU+costV <= budget; costV++ {
					if c0[costV] == NEG {
						continue
					}
					if valU+c0[costV] > newBase[costU+costV] {
						newBase[costU+costV] = valU + c0[costV]
					}
				}
			}
			base = newBase
		}

		// Now base[c] is the profit for "u does NOT buy" (regardless of whether parent bought).
		// dp0_base = base (case parent NOT buy and u not buy)
		// dp1_base = base (case parent DID buy and u not buy) -> children still see parent NOT bought because u didn't buy.

		// Build buy arrays: we consider "u buys" option — when u buys, children see parent bought, so we must merge children.dp1
		// buy0 : u buys under parent NOT buy (cost = present[u], profit = future - present)
		buy0 := make([]int, budget+1)
		for i := 0; i <= budget; i++ {
			buy0[i] = NEG
		}
		costU0 := present[u]
		profitU0 := future[u] - present[u]
		if costU0 <= budget {
			buy0[costU0] = profitU0
		}
		for idx := 0; idx < len(children); idx++ {
			c1 := childDP1[idx]
			newBuy := make([]int, budget+1)
			for i := 0; i <= budget; i++ {
				newBuy[i] = NEG
			}
			for costCur := 0; costCur <= budget; costCur++ {
				if buy0[costCur] == NEG {
					continue
				}
				valCur := buy0[costCur]
				for costV := 0; costCur+costV <= budget; costV++ {
					if c1[costV] == NEG {
						continue
					}
					if valCur+c1[costV] > newBuy[costCur+costV] {
						newBuy[costCur+costV] = valCur + c1[costV]
					}
				}
			}
			buy0 = newBuy
		}

		// buy1 : u buys under parent DID buy (discount price)
		buy1 := make([]int, budget+1)
		for i := 0; i <= budget; i++ {
			buy1[i] = NEG
		}
		costU1 := present[u] / 2
		profitU1 := future[u] - costU1
		if costU1 <= budget {
			buy1[costU1] = profitU1
		}
		for idx := 0; idx < len(children); idx++ {
			c1 := childDP1[idx]
			newBuy := make([]int, budget+1)
			for i := 0; i <= budget; i++ {
				newBuy[i] = NEG
			}
			for costCur := 0; costCur <= budget; costCur++ {
				if buy1[costCur] == NEG {
					continue
				}
				valCur := buy1[costCur]
				for costV := 0; costCur+costV <= budget; costV++ {
					if c1[costV] == NEG {
						continue
					}
					if valCur+c1[costV] > newBuy[costCur+costV] {
						newBuy[costCur+costV] = valCur + c1[costV]
					}
				}
			}
			buy1 = newBuy
		}

		// Compose final dp0 and dp1:
		dp0 := make([]int, budget+1)
		dp1 := make([]int, budget+1)
		for c := 0; c <= budget; c++ {
			dp0[c] = NEG
			dp1[c] = NEG
			// parent NOT buy, u NOT buy -> base[c]
			if base[c] != NEG {
				dp0[c] = base[c]
			}
			// parent NOT buy, u buys -> buy0[c]
			if buy0[c] != NEG && buy0[c] > dp0[c] {
				dp0[c] = buy0[c]
			}
			// parent DID buy, u NOT buy -> base[c] (u doesn't buy, children see parent NOT bought)
			if base[c] != NEG {
				dp1[c] = base[c]
			}
			// parent DID buy, u buys at discount -> buy1[c]
			if buy1[c] != NEG && buy1[c] > dp1[c] {
				dp1[c] = buy1[c]
			}
		}

		return dp0, dp1
	}

	dp0Root, _ := dfs(0)
	ans := 0
	for c := 0; c <= budget; c++ {
		if dp0Root[c] > ans {
			ans = dp0Root[c]
		}
	}
	return ans
}
