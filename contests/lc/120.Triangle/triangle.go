// Package triangle implements helpers for working with the classic
// "Triangle" minimum-path-sum dynamic programming problem.
package triangle

// minimumTotal returns the smallest possible sum of a path that starts at the
// apex of the triangle and, moving only to adjacent numbers on the row below,
// reaches the base. The implementation runs bottom-up, reusing the last row as
// a rolling buffer so it achieves O(n^2) time with O(n) additional space.
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	f := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			f[j] = min(f[j], f[j+1]) + triangle[i][j]
		}
	}
	return f[0]
}
