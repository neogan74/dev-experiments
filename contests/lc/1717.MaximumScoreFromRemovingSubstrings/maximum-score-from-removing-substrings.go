package maximumscorefromremovingsubstrings

func maximumGain(s string, x int, y int) (ans int) {
	a, b := 'a', 'b'
	if x < y {
		x, y = y, x
		a, b = b, a
	}

	var cnt1, cnt2 int
	for _, c := range s {
		switch c {
		case a:
			cnt1++
		case b:
			if cnt1 > 0 {
				ans += x
				cnt1--
			} else {
				cnt2++
			}
		default:
			ans += min(cnt1, cnt2) * y
			cnt1, cnt2 = 0, 0
		}
	}
	ans += min(cnt1, cnt2) * y
	return
}
