package main

// TreeNode is the standard LeetCode binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	level := 0
	bestLevel := 1
	bestSum := root.Val

	for len(queue) > 0 {
		level++
		levelSum := 0
		next := make([]*TreeNode, 0, len(queue)*2)
		for _, node := range queue {
			levelSum += node.Val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		if levelSum > bestSum || level == 1 {
			bestSum = levelSum
			bestLevel = level
		}
		queue = next
	}

	return bestLevel
}
