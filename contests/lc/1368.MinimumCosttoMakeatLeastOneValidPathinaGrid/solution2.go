package minimumcosttomakeatleastonevalidpathinagrid

func minCost2(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	size := m * n

	// Directions: 1=right, 2=left, 3=down, 4=up
	dirs := [][3]int{
		{0, 1, 1},
		{0, -1, 2},
		{1, 0, 3},
		{-1, 0, 4},
	}

	// ---- Fast custom deque (ring buffer) ----
	// Stores ints (flattened cell index)
	deque := make([]int, size*2)
	head, tail := size, size // start in middle to avoid wrap issues

	pushFront := func(x int) {
		head--
		deque[head] = x
	}
	pushBack := func(x int) {
		deque[tail] = x
		tail++
	}
	popFront := func() int {
		x := deque[head]
		head++
		return x
	}
	empty := func() bool {
		return head == tail
	}

	// ---- dist array ----
	const INF = int(1e9)
	dist := make([]int, size)
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0

	pushFront(0) // starting cell

	for !empty() {
		idx := popFront()
		r := idx / n
		c := idx % n
		d := dist[idx]

		// explore neighbors
		for _, v := range dirs {
			dr, dc, arrow := v[0], v[1], v[2]
			nr, nc := r+dr, c+dc

			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}

			nidx := nr*n + nc
			cost := 0
			if grid[r][c] != arrow {
				cost = 1
			}

			nd := d + cost
			if nd < dist[nidx] {
				dist[nidx] = nd
				if cost == 0 {
					pushFront(nidx)
				} else {
					pushBack(nidx)
				}
			}
		}
	}

	return dist[size-1]
}
