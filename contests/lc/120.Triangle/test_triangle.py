import pytest

from triangle import Solution


@pytest.mark.parametrize(
    "triangle, expected",
    [
        ([[2], [3, 4], [6, 5, 7], [4, 1, 8, 3]], 11),
        ([[5]], 5),
        ([[1], [-2, -3], [1, -1, -3]], -5),
    ],
)
def test_minimum_total(triangle, expected):
    assert Solution().minimumTotal(triangle) == expected
