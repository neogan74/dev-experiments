from typing import Optional, Tuple


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val: int = 0, left: "Optional[TreeNode]" = None, right: "Optional[TreeNode]" = None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def subtreeWithAllDeepest(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        node, _ = self._deepest_subtree(root)
        return node

    def _deepest_subtree(self, node: Optional[TreeNode]) -> Tuple[Optional[TreeNode], int]:
        if node is None:
            return None, 0

        left_node, left_depth = self._deepest_subtree(node.left)
        right_node, right_depth = self._deepest_subtree(node.right)

        if left_depth > right_depth:
            return left_node, left_depth + 1
        if right_depth > left_depth:
            return right_node, right_depth + 1
        return node, left_depth + 1
