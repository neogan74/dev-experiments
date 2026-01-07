package main

const mod int64 = 1_000_000_007

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sums := make([]int64, 0, 1)
	total := subtreeSum(root, &sums)

	var best int64
	for _, s := range sums {
		product := s * (total - s)
		if product > best {
			best = product
		}
	}

	return int(best % mod)
}

func subtreeSum(node *TreeNode, sums *[]int64) int64 {
	if node == nil {
		return 0
	}

	left := subtreeSum(node.Left, sums)
	right := subtreeSum(node.Right, sums)
	current := left + right + int64(node.Val)
	*sums = append(*sums, current)
	return current
}
