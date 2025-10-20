from max_num_oper import Solution

def test_examples():
    assert Solution().maxDistinctElements([1,2,2,3,3,4], 2) == 6
    assert Solution().maxDistinctElements([1,1,1], 0) == 1

def test_mixed():
    # можно разнести дубликаты за счёт k
    assert Solution().maxDistinctElements([5,5,5,6], 1) == 4  # [4,5,6,6]-> по жадному выйдет 4 уникальных
    assert Solution().maxDistinctElements([0,0,0,0], 2) == 4

def test_edge_cases():
    assert Solution().maxDistinctElements([], 10) == 0
    assert Solution().maxDistinctElements([7], 0) == 1
    assert Solution().maxDistinctElements([-5,-5,-4], 1) == 3
