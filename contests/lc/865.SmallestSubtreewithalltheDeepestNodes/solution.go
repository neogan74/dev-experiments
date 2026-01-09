package smallestsubtreewithallthedeepestnodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node, _ := deepestSubtree(root)
	return node
}

func deepestSubtree(node *TreeNode) (*TreeNode, int) {
	if node == nil {
		return nil, 0
	}

	leftNode, leftDepth := deepestSubtree(node.Left)
	rightNode, rightDepth := deepestSubtree(node.Right)

	if leftDepth > rightDepth {
		return leftNode, leftDepth + 1
	}
	if rightDepth > leftDepth {
		return rightNode, rightDepth + 1
	}
	return node, leftDepth + 1
}
