package main

import (
	"fmt"
)

// Определение структуры TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// helper возвращает пару (LCA, глубина) для поддерева с корнем в node.
func helper(node *TreeNode) (*TreeNode, int) {
	if node == nil {
		return nil, 0
	}

	leftLCA, leftDepth := helper(node.Left)
	rightLCA, rightDepth := helper(node.Right)

	if leftDepth > rightDepth {
		return leftLCA, leftDepth + 1
	} else if rightDepth > leftDepth {
		return rightLCA, rightDepth + 1
	} else {
		// Глубины равны – текущий узел является LCA.
		return node, leftDepth + 1
	}
}

// lcaDeepestLeaves возвращает наименьшего общего предка самых глубоких листьев.
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	lca, _ := helper(root)
	return lca
}

func main() {
	// Пример 1:
	// Построим дерево:
	//         3
	//        / \
	//       5   1
	//      / \   \
	//     6   2   8
	//        / \
	//       7   4
	//
	// Глубокие листья: 6, 7, 4, 8. Наименьший общий предок для 7 и 4 – узел 2, а для всей группы – узел 3 или 2?
	// Здесь самый глубокий уровень – узлы 7, 4, 8. Их LCA будет узел 3, если смотреть на всю группу.
	//
	// Для демонстрации построим дерево:
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	result := lcaDeepestLeaves(root)
	fmt.Printf("LCA самых глубоких листьев: %d\n", result.Val)
}
