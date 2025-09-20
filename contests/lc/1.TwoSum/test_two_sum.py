import pytest

from two_sum import Solution


@pytest.mark.parametrize(
    "nums, target, expected",
    [
        ([2, 7, 11, 15], 9, [0, 1]),
        ([3, 2, 4], 6, [1, 2]),
        ([3, 3], 6, [0, 1]),
        ([-1, 0, 1, 2], 1, [1, 2]),
    ],
)
def test_two_sum_variants(nums, target, expected):
    solution = Solution()
    assert solution.twoSum(nums, target) == expected
    assert solution.twoSum2(nums, target) == expected
