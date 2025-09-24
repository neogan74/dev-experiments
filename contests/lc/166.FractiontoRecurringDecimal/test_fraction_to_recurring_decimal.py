import importlib.util
from pathlib import Path

import pytest


@pytest.fixture(scope="module")
def solution_cls():
    module_path = Path(__file__).with_name("fraction-to-recurring-decimal.py")
    spec = importlib.util.spec_from_file_location(
        "fraction_to_recurring_decimal", module_path
    )
    module = importlib.util.module_from_spec(spec)
    assert spec.loader is not None
    spec.loader.exec_module(module)
    return module.Solution


@pytest.mark.parametrize(
    "numerator, denominator, expected",
    [
        (1, 2, "0.5"),
        (2, 1, "2"),
        (4, 333, "0.(012)"),
        (1, 6, "0.1(6)"),
        (-50, 8, "-6.25"),
        (7, -12, "-0.58(3)"),
        (1, 333, "0.(003)"),
        (1, 5, "0.2"),
        (0, -5, "0"),
        (1, -2147483648, "-0.0000000004656612873077392578125"),
    ],
)
def test_fraction_to_decimal(solution_cls, numerator, denominator, expected):
    solution = solution_cls()
    assert solution.fractionToDecimal(numerator, denominator) == expected


@pytest.mark.parametrize(
    "numerator, denominator, expected",
    [
        (1, 7, "0.(142857)"),
        (22, 7, "3.(142857)"),
        (1, 90, "0.0(1)"),
    ],
)
def test_fraction_to_decimal_repeating_sequences(solution_cls, numerator, denominator, expected):
    solution = solution_cls()
    assert solution.fractionToDecimal(numerator, denominator) == expected
