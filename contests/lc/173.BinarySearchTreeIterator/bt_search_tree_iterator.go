package _173_binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{}
	it.pushLeft(root)
	return it
}

func (it *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
}

func (it *BSTIterator) Next() int {
	n := len(it.stack)
	node := it.stack[n-1]
	it.stack = it.stack[:n-1]

	if node.Right != nil {
		it.pushLeft(node.Right)
	}

	return node.Val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
