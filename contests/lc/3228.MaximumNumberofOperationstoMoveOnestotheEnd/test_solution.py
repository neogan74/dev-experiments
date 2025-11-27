import itertools
from functools import lru_cache

import pytest

from solution import Solution


def brute_force_max_operations(s: str) -> int:
    """Compute the maximum number of operations by exploring every sequence."""

    @lru_cache(maxsize=None)
    def dfs(state: str) -> int:
        best = 0
        for next_state in enumerate_operations(state):
            best = max(best, 1 + dfs(next_state))
        return best

    return dfs(s)


def enumerate_operations(state: str):
    arr = list(state)
    n = len(arr)
    for i in range(n - 1):
        if state[i] == "1" and state[i + 1] == "0":
            step = arr.copy()
            j = i
            while j + 1 < n and step[j + 1] == "0":
                step[j], step[j + 1] = step[j + 1], step[j]
                j += 1
            yield "".join(step)


@pytest.mark.parametrize(
    ("s", "expected"),
    [
        ("1001101", 4),
        ("00111", 0),
        ("10", 1),
        ("1", 0),
        ("010010", 3),
    ],
)
def test_known_cases(s, expected):
    assert Solution().maxOperations(s) == expected


def test_matches_brute_force_for_small_strings():
    sol = Solution()
    for n in range(1, 7):
        for bits in itertools.product("01", repeat=n):
            s = "".join(bits)
            assert sol.maxOperations(s) == brute_force_max_operations(s)
