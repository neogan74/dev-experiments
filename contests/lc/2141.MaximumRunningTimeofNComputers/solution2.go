package maximumrunningtimeofncomputers

func maxRunTime2(n int, batteries []int) int64 {
	bb := batteries
	sum := 0
	for _, x := range bb {
		sum += x
	}

	canPower := func(x int) bool {
		t := 0
		goal := x * n
		for _, y := range bb {
			t += min(x, y)
			if t >= goal {
				return true
			}
		}
		return false
	}
	l, r := 1, sum/n
	for l <= r {
		m := (l + r) >> 1
		if canPower(m) {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return int64(r)
}
