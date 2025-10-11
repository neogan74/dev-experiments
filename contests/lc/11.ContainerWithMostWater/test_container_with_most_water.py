import importlib.util
from pathlib import Path
from typing import Callable, List

import pytest


def load_max_area() -> Callable[[List[int]], int]:
    module_path = Path(__file__).resolve().parent / "container-with-most-water.py"
    spec = importlib.util.spec_from_file_location("container_with_most_water", module_path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module.MaxArea


@pytest.fixture(scope="module")
def max_area() -> Callable[[List[int]], int]:
    return load_max_area()


@pytest.mark.parametrize(
    "heights, expected",
    [
        ([1, 8, 6, 2, 5, 4, 8, 3, 7], 49),
        ([1, 1], 1),
        ([4, 3, 2, 1, 4], 16),
        ([1, 2, 1], 2),
        ([2, 3, 4, 5, 18, 17, 6], 17),
    ],
)
def test_max_area_returns_expected_values(max_area, heights, expected):
    assert max_area(heights) == expected


def test_max_area_handles_single_peak_between_lower_values(max_area):
    heights = [2, 3, 10, 5, 7, 8, 9]
    assert max_area(heights) == 36  # Between heights[2]=10 and heights[6]=9 with width 4
