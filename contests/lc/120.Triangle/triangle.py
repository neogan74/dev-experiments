from typing import List


class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        if not triangle:
            return 0

        dp = triangle[-1][:]
        for i in range(len(triangle) - 2, -1, -1):
            row = triangle[i]
            for j in range(len(row)):
                dp[j] = min(dp[j], dp[j + 1]) + row[j]
        return dp[0]
