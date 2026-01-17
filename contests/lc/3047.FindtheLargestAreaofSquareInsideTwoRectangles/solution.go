package main

func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
	maxSide := int64(0)
	n := len(bottomLeft)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			left := maxInt(bottomLeft[i][0], bottomLeft[j][0])
			right := minInt(topRight[i][0], topRight[j][0])
			bottom := maxInt(bottomLeft[i][1], bottomLeft[j][1])
			top := minInt(topRight[i][1], topRight[j][1])

			width := right - left
			height := top - bottom
			if width <= 0 || height <= 0 {
				continue
			}

			side := int64(width)
			if height < width {
				side = int64(height)
			}
			if side > maxSide {
				maxSide = side
			}
		}
	}

	return maxSide * maxSide
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
