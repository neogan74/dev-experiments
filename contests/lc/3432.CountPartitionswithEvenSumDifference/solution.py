from typing import List


def countPartitions(nums: List[int]) -> int:
    """
    Return the number of partitions where the difference between left and right
    subarray sums is even.

    The difference is 2*prefix - total, so its parity matches the total sum.
    If the total is even, every split (n-1 of them) works; if odd, none do.
    """
    total = sum(nums)
    return len(nums) - 1 if total % 2 == 0 else 0
