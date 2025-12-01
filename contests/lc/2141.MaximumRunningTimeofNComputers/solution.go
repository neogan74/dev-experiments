package maximumrunningtimeofncomputers

// maxRunTime returns the maximum minutes all n computers can run simultaneously.
// Uses binary search on time with a feasibility check based on total usable energy.
// Example: n=2, batteries=[3,3,3]
// total=9 -> hi=4. mid=2 (feasible), mid=3 (feasible), mid=4 (feasible) => answer=4.
//
// Flow:
// 1) Search time in [0, total/n].
// 2) time t is feasible when sum(min(battery, t)) >= n*t.
// 3) Keep the largest feasible t.
func maxRunTime(n int, batteries []int) int64 {
	var total int64
	for _, b := range batteries {
		total += int64(b)
	}

	lo, hi := int64(0), total/int64(n)

	canRun := func(t int64) bool {
		var energy int64
		for _, b := range batteries {
			if int64(b) >= t {
				energy += t
			} else {
				energy += int64(b)
			}
		}
		return energy >= t*int64(n)
	}

	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canRun(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}
