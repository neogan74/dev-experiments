import importlib.util
import itertools
from pathlib import Path

import pytest

_MODULE_PATH = Path(__file__).with_name("valid-triangle-number.py")
_spec = importlib.util.spec_from_file_location("valid_triangle_number", _MODULE_PATH)
_module = importlib.util.module_from_spec(_spec)
_spec.loader.exec_module(_module)
Solution = _module.Solution


def brute_triangle_count(nums):
    count = 0
    for i, j, k in itertools.combinations(range(len(nums)), 3):
        a, b, c = sorted((nums[i], nums[j], nums[k]))
        if a + b > c:
            count += 1
    return count


@pytest.mark.parametrize(
    "nums, expected",
    [
        ([2, 2, 3, 4], 3),
        ([4, 2, 3, 4], 4),
        ([0, 0, 0], 0),
        ([1, 1, 3, 4], 0),
        ([5, 5, 5, 5], 4),
        ([1, 2], 0),
    ],
)
def test_examples_and_edge_cases(nums, expected):
    assert Solution().triangleNumber(nums.copy()) == expected


@pytest.mark.parametrize(
    "nums",
    [
        [],
        [1],
        [2, 2],
        [2, 3, 4],
        [2, 10, 12, 15],
        [0, 2, 0, 3, 2, 4],
        [10, 21, 22, 100, 101, 200, 300],
    ],
)
def test_triangle_number_matches_bruteforce(nums):
    assert Solution().triangleNumber(nums.copy()) == brute_triangle_count(nums)
