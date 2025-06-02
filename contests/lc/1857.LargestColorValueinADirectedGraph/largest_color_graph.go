package _1857_largest_color_graph

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	graph := make([][]int, n)
	inDegree := make([]int, n)

	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		inDegree[to]++
	}

	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// dp[node][color] — сколько раз цвет встречается на пути к node
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	for i := 0; i < n; i++ {
		color := colors[i] - 'a'
		dp[i][color] = 1
	}

	visited := 0
	maxColorCount := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		visited++

		for _, neighbor := range graph[node] {
			for c := 0; c < 26; c++ {
				add := 0
				if int(colors[neighbor]-'a') == c {
					add = 1
				}
				if dp[neighbor][c] < dp[node][c]+add {
					dp[neighbor][c] = dp[node][c] + add
				}
			}
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}

		for c := 0; c < 26; c++ {
			if dp[node][c] > maxColorCount {
				maxColorCount = dp[node][c]
			}
		}
	}

	if visited < n {
		return -1 // Цикл обнаружен
	}
	return maxColorCount
}
