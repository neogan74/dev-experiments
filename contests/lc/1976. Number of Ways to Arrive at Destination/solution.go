package main

import (
	"container/heap"
	"fmt"
)

const mod int = 1e9 + 7

type Edge struct {
	to, time int
}

type Item struct {
	node, time int
}

type PQ []Item

func (pq PQ) Len() int            { return len(pq) }
func (pq PQ) Less(i, j int) bool  { return pq[i].time < pq[j].time }
func (pq PQ) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PQ) Push(x interface{}) { *pq = append(*pq, x.(Item)) }

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func countPaths(n int, roads [][]int) int {
	graph := make([][]Edge, n)
	for _, r := range roads {
		u, v, t := r[0], r[1], r[2]
		graph[u] = append(graph[u], Edge{v, t})
		graph[v] = append(graph[v], Edge{u, t})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = 1 << 60 // большое число (infinity)
	}
	ways := make([]int, n)
	dist[0] = 0
	ways[0] = 1

	pq := &PQ{}
	heap.Push(pq, Item{0, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		curNode, curTime := cur.node, cur.time

		if curTime > dist[curNode] {
			continue
		}

		for _, edge := range graph[curNode] {
			nextNode := edge.to
			nextTime := curTime + edge.time

			if nextTime < dist[nextNode] {
				dist[nextNode] = nextTime
				ways[nextNode] = ways[curNode]
				heap.Push(pq, Item{nextNode, nextTime})
			} else if nextTime == dist[nextNode] {
				ways[nextNode] = (ways[nextNode] + ways[curNode]) % mod
			}
		}
	}

	return ways[n-1]
}

func main() {
	fmt.Println(countPaths(7, [][]int{
		{0, 6, 7}, {0, 1, 2}, {1, 2, 3}, {1, 3, 3}, {6, 3, 3},
		{3, 5, 1}, {6, 5, 1}, {2, 5, 1}, {0, 4, 5}, {4, 6, 2},
	})) // 4

	fmt.Println(countPaths(2, [][]int{
		{1, 0, 10},
	})) // 1
}
