import pytest

from small_miss import find_smallest_integer


@pytest.mark.parametrize(
    "nums,value,expected",
    [
        ([1, -10, 7, 13, 6, 8], 5, 4),
        ([1, -10, 7, 13, 6, 8], 7, 2),
        ([0, 1, 2], 1, 3),
        ([2, 2, 2], 2, 1),
        ([], 1, 0),
    ],
)
def test_find_smallest_integer(nums, value, expected):
    assert find_smallest_integer(nums, value) == expected
