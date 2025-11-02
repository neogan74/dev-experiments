import pytest

from solution import Solution


@pytest.mark.parametrize(
    "m, n, guards, walls, expected",
    [
        (
            4,
            6,
            [[0, 0], [1, 1], [2, 3]],
            [[0, 1], [2, 2], [1, 4]],
            7,
        ),
        (
            3,
            3,
            [[1, 1]],
            [[0, 1], [1, 0], [2, 1], [1, 2]],
            4,
        ),
        (
            5,
            5,
            [[0, 0], [4, 4]],
            [],
            9,
        ),
        (
            3,
            4,
            [[0, 0]],
            [[1, 0]],
            7,
        ),
    ],
)
def test_count_unguarded(m, n, guards, walls, expected):
    solution = Solution()
    assert solution.countUnguarded(m, n, guards, walls) == expected
