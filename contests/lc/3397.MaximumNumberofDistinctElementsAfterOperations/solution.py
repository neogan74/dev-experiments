from __future__ import annotations

from typing import List


def max_distinct_elements(nums: List[int], k: int) -> int:
    """Return the maximum number of distinct values achievable after operations."""
    if not nums:
        return 0
    sorted_nums = sorted(nums)
    occupied = -10**18
    distinct = 0
    for value in sorted_nums:
        candidate = max(occupied + 1, value - k)
        if candidate <= value + k:
            occupied = candidate
            distinct += 1
    return distinct


__all__ = ["max_distinct_elements"]
