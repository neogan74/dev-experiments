package main

import (
	"fmt"
)

// DSU - структура для поиска компоненты
type DSU struct {
	parent []int
	cost   []int
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	c := make([]int, n)
	for i := range p {
		p[i] = i
		c[i] = (1 << 30) - 1 // изначально все биты включены
	}
	return &DSU{parent: p, cost: c}
}

func (d *DSU) find(u int) int {
	if d.parent[u] != u {
		d.parent[u] = d.find(d.parent[u])
	}
	return d.parent[u]
}

func (d *DSU) union(u, v, w int) {
	pu, pv := d.find(u), d.find(v)
	if pu != pv {
		d.parent[pu] = pv
		d.cost[pv] &= d.cost[pu] & w
	} else {
		d.cost[pv] &= w
	}
}

func minimumCostWalk(n int, edges [][]int, query [][]int) []int {
	dsu := NewDSU(n)

	// объединяем вершины в компоненты
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		dsu.union(u, v, w)
	}

	res := make([]int, len(query))
	for i, q := range query {
		u, v := q[0], q[1]
		if dsu.find(u) != dsu.find(v) {
			res[i] = -1
		} else {
			res[i] = dsu.cost[dsu.find(u)]
		}
	}
	return res
}

func main() {
	// Примеры из условия задачи:
	fmt.Println(minimumCostWalk(5, [][]int{
		{0, 1, 7}, {1, 3, 7}, {1, 2, 1},
	}, [][]int{{0, 3}, {3, 4}})) // [1, -1]

	fmt.Println(minimumCostWalk(3, [][]int{
		{0, 2, 7}, {0, 1, 15}, {1, 2, 6}, {1, 2, 1},
	}, [][]int{{1, 2}})) // [0]
}
