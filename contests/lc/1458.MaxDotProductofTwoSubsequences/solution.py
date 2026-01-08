from typing import List


def maxDotProduct(nums1: List[int], nums2: List[int]) -> int:
    neg_inf = -10**18
    n = len(nums1)
    m = len(nums2)
    dp = [[neg_inf] * (m + 1) for _ in range(n + 1)]

    for i in range(1, n + 1):
        for j in range(1, m + 1):
            prod = nums1[i - 1] * nums2[j - 1]
            best = prod
            if dp[i - 1][j - 1] > 0:
                best = prod + dp[i - 1][j - 1]
            if dp[i - 1][j] > best:
                best = dp[i - 1][j]
            if dp[i][j - 1] > best:
                best = dp[i][j - 1]
            dp[i][j] = best

    return dp[n][m]
