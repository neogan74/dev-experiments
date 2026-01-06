import unittest

from solution import Solution, TreeNode


def build_tree(values):
    if not values or values[0] is None:
        return None
    nodes = [TreeNode(v) if v is not None else None for v in values]
    for i, node in enumerate(nodes):
        if node is None:
            continue
        li = 2 * i + 1
        ri = 2 * i + 2
        if li < len(nodes):
            node.left = nodes[li]
        if ri < len(nodes):
            node.right = nodes[ri]
    return nodes[0]


class TestMaxLevelSum(unittest.TestCase):
    def test_example1(self):
        root = build_tree([1, 7, 0, 7, -8, None, None])
        self.assertEqual(Solution().maxLevelSum(root), 2)

    def test_example2(self):
        root = build_tree([989, None, 10250, 98693, -89388, None, None, None, -32127])
        self.assertEqual(Solution().maxLevelSum(root), 2)

    def test_single_node(self):
        root = build_tree([-5])
        self.assertEqual(Solution().maxLevelSum(root), 1)


if __name__ == "__main__":
    unittest.main()
