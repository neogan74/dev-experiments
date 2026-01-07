from typing import Optional


class TreeNode:
    def __init__(self, val: int = 0, left: Optional["TreeNode"] = None, right: Optional["TreeNode"] = None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxProduct(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        sums = []
        total = self._subtree_sum(root, sums)
        best = 0
        for s in sums:
            product = s * (total - s)
            if product > best:
                best = product
        return best % 1_000_000_007

    def _subtree_sum(self, node: Optional[TreeNode], sums: list[int]) -> int:
        if node is None:
            return 0

        left = self._subtree_sum(node.left, sums)
        right = self._subtree_sum(node.right, sums)
        current = left + right + node.val
        sums.append(current)
        return current
