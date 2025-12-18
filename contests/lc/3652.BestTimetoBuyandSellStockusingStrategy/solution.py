from typing import List

class Solution:
    def maxProfit(self, prices: List[int], strategy: List[int], k: int) -> int:
        """Return the maximum profit after applying at most one allowed modification."""
        n = len(prices)
        if n == 0 or k == 0:
            return 0

        half = k // 2

        prefix_price = [0] * (n + 1)
        prefix_prod = [0] * (n + 1)  # strategy[i] * prices[i]

        for i, (price, action) in enumerate(zip(prices, strategy)):
            prefix_price[i + 1] = prefix_price[i] + price
            prefix_prod[i + 1] = prefix_prod[i] + price * action

        base = prefix_prod[-1]
        best_delta = 0

        for start in range(0, n - k + 1):
            mid = start + half
            end = start + k

            first_half = prefix_prod[mid] - prefix_prod[start]  # strategy * price
            second_price = prefix_price[end] - prefix_price[mid]  # price
            second_prod = prefix_prod[end] - prefix_prod[mid]  # strategy * price
            delta = -first_half + (second_price - second_prod)  # profit gain if modified

            if delta > best_delta:
                best_delta = delta

        return base + best_delta

    def maxProfit2(self, prices: List[int], strategy: List[int], k: int) -> int:
        n = len(prices)
        m = k // 2

        # Baseline profit
        A = [strategy[i] * prices[i] for i in range(n)]
        base = sum(A)

        # Sliding window to compute best delta for one modification
        sumA = sum(A[:k])
        sumP2 = sum(prices[m:k])  # prices of the second half for window starting at 0
        bestDelta = sumP2 - sumA

        for l in range(1, n - k + 1):
            # Update sum of A over the whole window [l, l+k-1]
            sumA += A[l + k - 1] - A[l - 1]
            # Update sum of prices over the second half [l+m, l+k-1]
            sumP2 += prices[l + k - 1] - prices[l + m - 1]
            cur = sumP2 - sumA
            if cur > bestDelta:
                bestDelta = cur

        if bestDelta < 0:
            bestDelta = 0

        return base + bestDelta