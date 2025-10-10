import random

import pytest

from find_the_min_amount_of_time import minTime


def no_wait_baseline(skill, mana):
    n, m = len(skill), len(mana)
    if n == 0 or m == 0:
        return 0
    times = [[s * ma for ma in mana] for s in skill]
    column_prefix = [[0] * n for _ in range(m)]
    for j in range(m):
        total = 0
        for i in range(n):
            total += times[i][j]
            column_prefix[j][i] = total
    starts = [0] * m
    for j in range(1, m):
        best = starts[j - 1] + times[0][j - 1]
        for i in range(1, n):
            candidate = starts[j - 1] + column_prefix[j - 1][i] - column_prefix[j][i - 1]
            if candidate > best:
                best = candidate
        starts[j] = best
    return starts[-1] + column_prefix[-1][-1]


def test_examples_from_description():
    assert minTime([1, 5, 2, 4], [5, 1, 4, 2]) == 110
    assert minTime([1, 1, 1], [1, 1, 1]) == 5
    assert minTime([1, 2, 3, 4], [1, 2]) == 21


def test_single_and_multi_wizard_potions():
    assert minTime([1], [5]) == 5
    assert minTime([1, 2], [3]) == 9
    assert minTime([2, 3], [4, 1]) == no_wait_baseline([2, 3], [4, 1])


@pytest.mark.parametrize(
    "skill,mana",
    [
        ([1, 2, 3], [1, 2]),
        ([2, 2], [3, 1, 4]),
        ([5], [1, 2, 3, 4]),
        ([3, 1, 4], [2, 2, 1]),
    ],
)
def test_matches_baseline_for_known_cases(skill, mana):
    assert minTime(skill, mana) == no_wait_baseline(skill, mana)


def test_random_small_instances_match_baseline():
    rng = random.Random(0)
    for _ in range(50):
        n = rng.randint(1, 4)
        m = rng.randint(1, 4)
        skill = [rng.randint(1, 5) for _ in range(n)]
        mana = [rng.randint(1, 5) for _ in range(m)]
        assert minTime(skill, mana) == no_wait_baseline(skill, mana)


def test_zero_length_inputs_returns_zero():
    assert minTime([], []) == 0
    assert minTime([1, 2, 3], []) == 0
    assert minTime([], [1, 2, 3]) == 0
