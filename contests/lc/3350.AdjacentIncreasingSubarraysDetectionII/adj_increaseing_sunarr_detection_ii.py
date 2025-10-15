from typing import List

class Solution:
    def maxIncreasingSubarrays(self, nums: List[int]) -> int:
        ans = pre = cur = 0
        for i, x in enumerate(nums):
            cur += 1
            if i == len(nums) - 1 or x >= nums[i + 1]:
                ans = max(ans, cur // 2, min(pre,cur))
                pre, cur = cur, 0
        return ans

if __name__ == "__main__":
    s = Solution()
    print(s.maxIncreasingSubarrays([2, 5, 7, 8, 9, 2, 3, 4, 3, 1])) # expected 3