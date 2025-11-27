package maximizetheminimumpoweredcity


func maxPower(stations []int, r int, k int) int64 {
	n := len(stations)
	d := make([]int, n+1)
	s := make([]int, n+1)
	for i, v := range stations {
		left, right := max(0, i-r), min(i+r, n-1)
		d[left] += v
		d[right+1] -= v
	}
	s[0] = d[0]
	for i := 1; i < n+1; i++ {
		s[i] = s[i-1] + d[i]
	}
	check := func(x, k int) bool {
		d := make([]int, n+1)
		t := 0
		for i := range stations {
			t += d[i]
			dist := x - (s[i] + t)
			if dist > 0 {
				if k < dist {
					return false
				}
				k -= dist
				j := min(i+r, n-1)
				left, right := max(0, j-r), min(j+r, n-1)
				d[left] += dist
				d[right+1] -= dist
				t += dist
			}
		}
		return true
	}
	left, right := 0, 1<<40
	for left < right {
		mid := (left + right + 1) >> 1
		if check(mid, k) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return int64(left)
}


func maxPower2(stations []int, r int, k int) int64 {
	n := len(stations)
	cnt := make([]int64, n+1)
	for i := 0; i < n; i++ {
		left := max(0, i-r)
		right := min(n, i+r+1)
		cnt[left] += int64(stations[i])
		cnt[right] -= int64(stations[i])
	}

	minVal := int64(stations[0])
	sumTotal := int64(0)
	for _, s := range stations {
		if int64(s) < minVal {
			minVal = int64(s)
		}
		sumTotal += int64(s)
	}

	lo, hi := minVal, sumTotal+int64(k)
	var res int64 = 0

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if check(cnt, mid, r, k) {
			res = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return res
}

func check(cnt []int64, val int64, r int, k int) bool {
	n := len(cnt) - 1
	diff := make([]int64, len(cnt))
	copy(diff, cnt)
	var sum int64 = 0
	remaining := int64(k)

	for i := 0; i < n; i++ {
		sum += diff[i]
		if sum < val {
			add := val - sum
			if remaining < add {
				return false
			}
			remaining -= add
			end := min(n, i+2*r+1)
			diff[end] -= add
			sum += add
		}
	}

	return true
}