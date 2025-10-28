from typing import List


class Solution:
    def countValidSelections(self, nums: List[int]) -> int:
        total = sum(nums)
        prefix = 0
        ans = 0

        for value in nums:
            if value != 0:
                prefix += value
                continue

            if prefix * 2 == total:
                ans += 2
            elif abs(prefix * 2 - total) <= 1:
                ans += 1

        return ans

if __name__ == "__main__":
    s = Solution()
    print(s.countValidSelections([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]))