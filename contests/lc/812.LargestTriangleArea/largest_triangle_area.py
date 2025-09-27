from __future__ import annotations

from itertools import combinations
from typing import Iterable, List, Sequence


def triangleArea(a: Sequence[int], b: Sequence[int], c: Sequence[int]) -> float:
    """Return area of triangle formed by three 2D points."""
    (x1, y1), (x2, y2), (x3, y3) = a, b, c
    return abs(x1 * (y2 - y3) + x2 * (y3 - y1) + x3 * (y1 - y2)) / 2.0


def largestTriangleArea(points: Iterable[Sequence[int]]) -> float:
    """Compute the maximum triangle area from the given points."""
    max_area = 0.0
    # Convert to list once so we can iterate combinations multiple times.
    pts: List[Sequence[int]] = list(points)
    for a, b, c in combinations(pts, 3):
        area = triangleArea(a, b, c)
        if area > max_area:
            max_area = area
    return max_area


__all__ = ["largestTriangleArea", "triangleArea"]
