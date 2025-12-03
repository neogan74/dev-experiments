package lengthoflongestvshapeddiagonalsegment

import "testing"

var (
	exampleGrid = [][]int{
		{2, 2, 1, 2, 2},
		{2, 0, 2, 2, 0},
		{2, 0, 1, 1, 0},
		{1, 0, 2, 2, 2},
		{2, 0, 0, 2, 2},
	}
	largeGrid = buildDiagonalGrid(200, 200)
)

func buildDiagonalGrid(rows, cols int) [][]int {
	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}
	diag := min(rows, cols)
	for i := 0; i < diag; i++ {
		switch i {
		case 0:
			grid[i][i] = 1
		case 1:
			grid[i][i] = 2
		default:
			if i%2 == 0 {
				grid[i][i] = 0
			} else {
				grid[i][i] = 2
			}
		}
	}
	return grid
}

func Benchmark_lenOfVDiagonalSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = lenOfVDiagonal(exampleGrid)
	}
}

func Benchmark_lenOfVDiagonalLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = lenOfVDiagonal(largeGrid)
	}
}
