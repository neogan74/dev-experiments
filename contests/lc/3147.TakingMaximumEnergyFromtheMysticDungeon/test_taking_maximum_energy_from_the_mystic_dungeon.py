import pytest

from taking_maximum_energy_from_the_mystic_dungeon import max_energy


@pytest.mark.parametrize(
    "energy,k,expected",
    [
        ([5, 2, -10, -5, 1], 3, 3),
        ([-2, -3, -1], 2, -1),
        ([4, -1, 2, 1, -5, 4], 2, 5),  # start at index 0 -> 4 + 2 - 5 = 1; index 1 -> -1 + 1 + 4 = 4; index 2 -> 2 -5 = -3; index 3 -> 1; index 4 -> -5; index 5 -> 4
        ([1], 1, 1),
    ],
)
def test_max_energy(energy, k, expected):
    assert max_energy(energy, k) == expected


def test_large_k():
    energy = [3, -2, 7]
    assert max_energy(energy, 2) == 10  # start at index 0 -> 3 + 7 = 10
