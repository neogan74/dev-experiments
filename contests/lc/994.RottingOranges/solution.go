package main

import "fmt"

type Point struct {
	x, y int
}

func orangesRotting(grid [][]int) int {
	dirs := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := []Point{}

	rows, cols := len(grid), len(grid[0])
	fresh := 0

	// Добавляем гнилые апельсины в очередь и считаем свежие
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			switch grid[i][j] {
			case 2:
				queue = append(queue, Point{i, j})
			case 1:
				fresh++
			}
		}
	}

	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			p := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nx, ny := p.x+d.x, p.y+d.y
				if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == 1 {
					grid[nx][ny] = 2
					fresh--
					queue = append(queue, Point{nx, ny})
				}
			}
		}
		minutes++
	}

	if fresh == 0 {
		return minutes
	}
	return -1
}

func main() {
	grid1 := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	fmt.Println(orangesRotting(grid1)) // ожидаем: 4

	grid2 := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}
	fmt.Println(orangesRotting(grid2)) // ожидаем: -1 (невозможно)

	grid3 := [][]int{{0, 2}}
	fmt.Println(orangesRotting(grid3)) // ожидаем: 0
}
