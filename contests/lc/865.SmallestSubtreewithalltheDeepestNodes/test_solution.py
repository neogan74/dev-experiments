from solution import Solution, TreeNode


def build_tree(values):
    if not values or values[0] is None:
        return None

    root = TreeNode(values[0])
    queue = [root]
    index = 1

    while index < len(values) and queue:
        node = queue.pop(0)

        if index < len(values):
            left_val = values[index]
            if left_val is not None:
                node.left = TreeNode(left_val)
                queue.append(node.left)
            index += 1

        if index < len(values):
            right_val = values[index]
            if right_val is not None:
                node.right = TreeNode(right_val)
                queue.append(node.right)
            index += 1

    return root


def tree_to_list(root):
    if root is None:
        return []

    result = []
    queue = [root]
    while queue:
        node = queue.pop(0)
        if node is None:
            result.append(None)
            continue
        result.append(node.val)
        queue.append(node.left)
        queue.append(node.right)

    while result and result[-1] is None:
        result.pop()

    return result


def test_subtree_with_all_deepest_example1():
    root = build_tree([3, 5, 1, 6, 2, 0, 8, None, None, 7, 4])
    result = Solution().subtreeWithAllDeepest(root)
    assert tree_to_list(result) == [2, 7, 4]


def test_subtree_with_all_deepest_example2():
    root = build_tree([1])
    result = Solution().subtreeWithAllDeepest(root)
    assert tree_to_list(result) == [1]


def test_subtree_with_all_deepest_example3():
    root = build_tree([0, 1, 3, None, 2])
    result = Solution().subtreeWithAllDeepest(root)
    assert tree_to_list(result) == [2]


def test_subtree_with_all_deepest_two_deepest_siblings():
    root = build_tree([1, 2, 3])
    result = Solution().subtreeWithAllDeepest(root)
    assert tree_to_list(result) == [1, 2, 3]
