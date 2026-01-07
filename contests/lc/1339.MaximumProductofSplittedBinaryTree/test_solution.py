from solution import Solution, TreeNode


def build_tree(values):
    if not values:
        return None
    if values[0] is None:
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


def test_max_product_example1():
    root = build_tree([1, 2, 3, 4, 5, 6])
    assert Solution().maxProduct(root) == 110


def test_max_product_example2():
    root = build_tree([1, None, 2, 3, 4, None, None, 5, 6])
    assert Solution().maxProduct(root) == 90
