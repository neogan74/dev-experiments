package _24_BinaryTreeMaximumPathSum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(dfs(node.Left), 0) // отрицательные не берём
		right := max(dfs(node.Right), 0)

		currentPath := node.Val + left + right
		maxSum = max(maxSum, currentPath)

		return node.Val + max(left, right) // возвращаем одну сторону
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
