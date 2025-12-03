package minimumcosttomakeatleastonevalidpathinagrid

import "container/heap"

type Point struct {
	x, y, cost int
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// Distance array to store minimum cost to reach each cell
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = 1e9 // Initialize with a very high value
		}
	}
	dist[0][0] = 0 // Start at (0, 0) with cost 0

	// Directions: Right (1), Left (2), Down (3), Up (4)
	directions := []struct{ dx, dy int }{
		{0, 1},  // Right
		{0, -1}, // Left
		{1, 0},  // Down
		{-1, 0}, // Up
	}

	// Min heap to process the cell with the smallest cost first
	h := &MinHeap{}
	heap.Push(h, Point{0, 0, 0}) // Push starting point (0, 0) with 0 cost

	// Dijkstra's algorithm
	for h.Len() > 0 {
		// Pop the cell with the smallest cost
		cell := heap.Pop(h).(Point)
		x, y, cost := cell.x, cell.y, cell.cost

		// If we have already visited this cell, continue
		if cost > dist[x][y] {
			continue
		}

		// Check all possible directions
		for i := 0; i < 4; i++ {
			// Calculate the next cell's coordinates
			newX, newY := x+directions[i].dx, y+directions[i].dy
			if newX < 0 || newY < 0 || newX >= m || newY >= n {
				continue
			}

			// Calculate the new cost to reach this cell
			newCost := cost
			if grid[x][y] != i+1 { // If the direction doesn't match, we need to change it
				newCost++
			}

			// If the new cost is better, update it and push to the heap
			if newCost < dist[newX][newY] {
				dist[newX][newY] = newCost
				heap.Push(h, Point{newX, newY, newCost})
			}
		}
	}

	// Return the minimum cost to reach the bottom-right cell
	return dist[m-1][n-1]
}
