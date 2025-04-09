package main

import "fmt"

func countCompleteComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	visited := make([]bool, n)
	var dfs func(int) (int, int)

	dfs = func(u int) (int, int) {
		visited[u] = true
		vertices, edges := 1, len(graph[u])

		for _, v := range graph[u] {
			if !visited[v] {
				vCount, eCount := dfs(v)
				vertices += vCount
				edges += eCount
			}
		}
		return vertices, edges
	}

	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			vertices, edges := dfs(i)
			edges /= 2 // Каждое ребро посчитано дважды
			if edges == vertices*(vertices-1)/2 {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}}))         // Ожидаем: 3
	fmt.Println(countCompleteComponents(6, [][]int{{0, 1}, {0, 2}, {1, 2}, {3, 4}, {3, 5}})) // Ожидаем: 1
}
