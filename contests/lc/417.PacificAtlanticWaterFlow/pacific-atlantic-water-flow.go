package _17_PacificAtlanticWaterFlow

func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}
	m, n := len(heights), len(heights[0])

	bfs := func(starts [][2]int) [][]bool {
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		q := make([][2]int, 0, len(starts))
		push := func(r, c int) { vis[r][c] = true; q = append(q, [2]int{r, c}) }
		for _, s := range starts {
			if !vis[s[0]][s[1]] {
				push(s[0], s[1])
			}
		}
		head := 0
		dir := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for head < len(q) {
			rc := q[head]
			head++
			r, c := rc[0], rc[1]
			h := heights[r][c]
			for _, d := range dir {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !vis[nr][nc] && heights[nr][nc] >= h {
					push(nr, nc)
				}
			}
		}
		return vis
	}

	pacStarts := make([][2]int, 0, m+n)
	for c := 0; c < n; c++ {
		pacStarts = append(pacStarts, [2]int{0, c})
	}
	for r := 0; r < m; r++ {
		pacStarts = append(pacStarts, [2]int{r, 0})
	}

	atlStarts := make([][2]int, 0, m+n)
	for c := 0; c < n; c++ {
		atlStarts = append(atlStarts, [2]int{m - 1, c})
	}
	for r := 0; r < m; r++ {
		atlStarts = append(atlStarts, [2]int{r, n - 1})
	}

	pac := bfs(pacStarts)
	atl := bfs(atlStarts)

	out := make([][]int, 0)
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pac[r][c] && atl[r][c] {
				out = append(out, []int{r, c})
			}
		}
	}
	return out
}
