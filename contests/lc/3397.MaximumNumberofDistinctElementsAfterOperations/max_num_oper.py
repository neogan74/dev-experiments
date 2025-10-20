
from typing import List

class Solution:
    def maxDistinctElements(self, nums: List[int], k: int) -> int:
        nums.sort()
        ans = 0
        occupied = -10**30  # "последнее занятое" значение
        for x in nums:
            # самое левое доступное место правее occupied
            place = max(occupied + 1, x - k)
            if place <= x + k:
                occupied = place
                ans += 1
        return ans
