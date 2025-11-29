from typing import List


class Solution:
    def minOperations(self, nums: List[int], k: int) -> int:
        """Return the minimal decrements needed to make sum(nums) divisible by k."""
        return sum(nums) % k
