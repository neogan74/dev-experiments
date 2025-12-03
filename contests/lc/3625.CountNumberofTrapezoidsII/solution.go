package countnumberoftrapezoidsii

// numberOfTrapezoids counts convex quadrilaterals with at least one pair of parallel sides.
// Idea:
//   - For any direction (slope), group all lines with that slope by their offset (perpendicular intercept).
//     Each pair of points on the same line contributes one to that line's pair count a_i = C(cnt_i, 2).
//     For this slope, the number of quadrilaterals whose parallel sides are on two different lines is
//     sum_{i<j} a_i * a_j = ((sum a_i)^2 - sum a_i^2) / 2.
//   - Every parallelogram is counted twice (once per slope of its two pairs of opposite sides), so subtract
//     the number of parallelograms. Parallelograms correspond to pairs of segments sharing the same midpoint.
//
// Complexity: O(n^2) time and O(n^2) space for maps over the O(n^2) point pairs.
func numberOfTrapezoids(points [][]int) int {
	type vec struct{ x, y int }

	gcd := func(a, b int) int {
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	// slope -> offset -> pair count on that line
	linePairs := map[vec]map[int]int{}
	// midpoint -> count of segments with that midpoint
	midCnt := map[vec]int{}

	for i, p1 := range points {
		x1, y1 := p1[0], p1[1]
		for _, p2 := range points[i+1:] {
			x2, y2 := p2[0], p2[1]
			dx, dy := x2-x1, y2-y1
			g := gcd(dx, dy)
			dx /= g
			dy /= g
			if dx < 0 || dx == 0 && dy < 0 { // canonical direction
				dx, dy = -dx, -dy
			}
			slope := vec{dx, dy}

			nx, ny := -dy, dx
			offset := nx*x1 + ny*y1
			if linePairs[slope] == nil {
				linePairs[slope] = map[int]int{}
			}
			linePairs[slope][offset]++

			mid := vec{x1 + x2, y1 + y2} // use doubled midpoint to keep integers
			midCnt[mid]++
		}
	}

	var parallelPairs int64
	for _, lines := range linePairs {
		var sumA, sumSq int64
		for _, a := range lines { // a = number of pairs on this line = C(cnt, 2)
			ai := int64(a)
			sumA += ai
			sumSq += ai * ai
		}
		parallelPairs += (sumA*sumA - sumSq) / 2
	}

	var parallelograms int64
	for _, k := range midCnt {
		if k > 1 {
			parallelograms += int64(k*(k-1)) / 2
		}
	}

	return int(parallelPairs - parallelograms)
}

func numberOfTrapezoids2(points [][]int) int {
	n := len(points)
	inf := 1e9 + 7
	slopeToIntercept := make(map[float64][]float64)
	midToSlope := make(map[float64][]float64)
	ans := 0

	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dx := x1 - x2
			dy := y1 - y2

			var k, b float64
			if x2 == x1 {
				k = inf
				b = float64(x1)
			} else {
				k = float64(y2-y1) / float64(x2-x1)
				b = float64(y1*dx-x1*dy) / float64(dx)
			}

			mid := float64((x1+x2)*10000 + (y1 + y2))
			slopeToIntercept[k] = append(slopeToIntercept[k], b)
			midToSlope[mid] = append(midToSlope[mid], k)
		}
	}

	for _, sti := range slopeToIntercept {
		if len(sti) == 1 {
			continue
		}

		cnt := make(map[float64]int)
		for _, bVal := range sti {
			cnt[bVal]++
		}

		totalSum := 0
		for _, count := range cnt {
			ans += totalSum * count
			totalSum += count
		}
	}

	for _, mts := range midToSlope {
		if len(mts) == 1 {
			continue
		}

		cnt := make(map[float64]int)
		for _, kVal := range mts {
			cnt[kVal]++
		}

		totalSum := 0
		for _, count := range cnt {
			ans -= totalSum * count
			totalSum += count
		}
	}

	return ans
}
