package sudokusolver

import "testing"

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solveSudoku(tt.args.board)
			isValidSolution(tt.args.board)
		})
	}
}

func isValidSolution(board [][]byte) bool {
	// Проверка строк
	for i := 0; i < 9; i++ {
		seen := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return false // Должны быть заполнены все клетки
			}
			if seen[board[i][j]] {
				return false
			}
			seen[board[i][j]] = true
		}
	}

	// Проверка столбцов
	for j := 0; j < 9; j++ {
		seen := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			if seen[board[i][j]] {
				return false
			}
			seen[board[i][j]] = true
		}
	}

	// Проверка 3x3 блоков
	for block := 0; block < 9; block++ {
		seen := make(map[byte]bool)
		r0 := (block / 3) * 3
		c0 := (block % 3) * 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				val := board[r0+i][c0+j]
				if seen[val] {
					return false
				}
				seen[val] = true
			}
		}
	}

	return true
}
