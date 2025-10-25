import pytest

from solution import next_beautiful_number


@pytest.mark.parametrize(
    "n, expected",
    [
        (1, 22),
        (1000, 1333),
        (3000, 3133),
        (0, 1),  # the smallest numerically balanced number is 1
        (22, 122),  # 122 has one '1' and two '2's
    ],
)
def test_next_beautiful_number_examples(n: int, expected: int) -> None:
    assert next_beautiful_number(n) == expected
