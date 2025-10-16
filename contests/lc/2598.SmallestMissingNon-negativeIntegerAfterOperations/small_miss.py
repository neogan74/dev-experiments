from collections import Counter
from typing import List


def find_smallest_integer(nums: List[int], value: int) -> int:
    """
    Return the maximum possible MEX after repeatedly adding or subtracting value from any element.

    The achievable residues modulo value form buckets; each bucket can supply consecutive integers
    until its count is exhausted.
    """
    remainder_counts = Counter((num % value + value) % value for num in nums)
    mex = 0
    while True:
        remainder = mex % value
        if remainder_counts[remainder] == 0:
            return mex
        remainder_counts[remainder] -= 1
        mex += 1
