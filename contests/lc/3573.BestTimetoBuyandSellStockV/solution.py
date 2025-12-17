from typing import List


class Solution:
    def maximumProfit(self, prices: List[int], k: int) -> int:
        if not prices or k == 0:
            return 0
        k = min(k, len(prices) // 2)

        NEG = -10**20
        rest = [NEG] * (k + 1)   # no open position
        long = [NEG] * (k + 1)   # currently long
        short = [NEG] * (k + 1)  # currently short
        rest[0] = 0

        for price in prices:
            next_rest = rest[:]   # carry over previous states
            next_long = long[:]
            next_short = short[:]

            for t in range(k + 1):
                if long[t] != NEG and t + 1 <= k:
                    next_rest[t + 1] = max(next_rest[t + 1], long[t] + price)
                if short[t] != NEG and t + 1 <= k:
                    next_rest[t + 1] = max(next_rest[t + 1], short[t] - price)
                if rest[t] != NEG:
                    next_long[t] = max(next_long[t], rest[t] - price)
                    next_short[t] = max(next_short[t], rest[t] + price)

            rest, long, short = next_rest, next_long, next_short

        return max(rest)
