package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Cell struct {
	val, x, y int
}

type MinHeap []Cell

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Cell)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	res := old[n-1]
	*h = old[:n-1]
	return res
}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	h := &MinHeap{}
	heap.Init(h)
	heap.Push(h, Cell{grid[0][0], 0, 0})
	visited[0][0] = true

	type Entry struct {
		threshold int
		count     int
	}
	record := []Entry{}
	count := 0
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		if len(record) == 0 || cell.val > record[len(record)-1].threshold {
			record = append(record, Entry{cell.val, count})
		}
		count++
		for _, d := range dirs {
			nx, ny := cell.x+d[0], cell.y+d[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx][ny] {
				visited[nx][ny] = true
				heap.Push(h, Cell{grid[nx][ny], nx, ny})
			}
		}
	}

	// Обрабатываем запросы
	result := make([]int, len(queries))
	for i, q := range queries {
		idx := sort.Search(len(record), func(i int) bool {
			return record[i].threshold >= q
		})
		if idx == 0 {
			result[i] = 0
		} else {
			result[i] = record[idx-1].count
		}
	}
	return result
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
	}
	queries := []int{4, 2, 6}
	fmt.Println("Result:", maxPoints(grid, queries)) // Ожидается: [6, 1, 6]
}
