package solution

func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	total := 0
	for i := 1; i < len(points); i++ {
		dx := abs(points[i][0] - points[i-1][0])
		dy := abs(points[i][1] - points[i-1][1])
		if dx > dy {
			total += dx
		} else {
			total += dy
		}
	}
	return total
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
