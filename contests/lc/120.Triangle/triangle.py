from typing import List


class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        """Return the minimum path sum from top to bottom of ``triangle``.

        The algorithm runs bottom-up, collapsing each row into a 1-D buffer so
        the time complexity stays O(n^2) while the additional space remains
        linear in the width of the triangle.
        """
        if not triangle:
            return 0

        dp = triangle[-1][:]
        for i in range(len(triangle) - 2, -1, -1):
            row = triangle[i]
            for j in range(len(row)):
                dp[j] = min(dp[j], dp[j + 1]) + row[j]
        return dp[0]
