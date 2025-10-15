import importlib.util
import pathlib

import pytest

# Load the solution module despite the hyphenated filename.
spec = importlib.util.spec_from_file_location(
    "solution",
    pathlib.Path(__file__).with_name("adjacent-increasing-subarrays-detection-i.py"),
)
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)
Solution = module.Solution


@pytest.mark.parametrize(
    "nums,k,expected",
    [
        ([2, 5, 7, 8, 9, 2, 3, 4, 3, 1], 3, True),
        ([1, 2, 3, 4, 4, 4, 4, 5, 6, 7], 5, False),
        ([1, 2, 3, 4, 5], 2, True),
        ([1, 2, 3, 0, 1, 2], 3, True),
        ([5, 4, 3, 2, 1, 0], 2, False),
    ],
)
def test_has_increasing_subarrays(nums, k, expected):
    assert Solution().hasIncreasingSubarrays(nums, k) == expected
