package _78_SwiminRisingWater

import "container/heap"

type node struct{ t, r, c int }
type minHeap []node

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(node)) }
func (h *minHeap) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func swimInWater(grid [][]int) int {
	n := len(grid)
	const INF = int(1e9)
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = INF
		}
	}
	dist[0][0] = grid[0][0]
	h := &minHeap{{grid[0][0], 0, 0}}
	heap.Init(h)
	dr := [5]int{1, 0, -1, 0, 1}

	for h.Len() > 0 {
		cur := heap.Pop(h).(node)
		if cur.r == n-1 && cur.c == n-1 {
			return cur.t
		}
		if cur.t != dist[cur.r][cur.c] {
			continue
		}
		for k := 0; k < 4; k++ {
			nr, nc := cur.r+dr[k], cur.c+dr[k+1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n {
				nt := cur.t
				if grid[nr][nc] > nt {
					nt = grid[nr][nc]
				}
				if nt < dist[nr][nc] {
					dist[nr][nc] = nt
					heap.Push(h, node{nt, nr, nc})
				}
			}
		}
	}
	return -1
}
