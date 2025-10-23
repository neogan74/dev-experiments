import pytest

from check_dig_after_oper import has_same_digits


@pytest.mark.parametrize(
    "s, expected",
    [
        ("3902", True),
        ("34789", False),
        ("112", True),  # reduces to "22"
        ("120", False),  # reduces to "30"
    ],
)
def test_has_same_digits_examples(s: str, expected: bool) -> None:
    assert has_same_digits(s) is expected
