package main

func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	dirs := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveNeighbors := 0
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n {
					if board[ni][nj] == 1 || board[ni][nj] == -1 {
						liveNeighbors++
					}
				}
			}

			if board[i][j] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				board[i][j] = -1 // умирает
			}
			if board[i][j] == 0 && liveNeighbors == 3 {
				board[i][j] = 2 // оживает
			}
		}
	}

	// нормализация
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 0
			} else if board[i][j] == 2 {
				board[i][j] = 1
			}
		}
	}
}
