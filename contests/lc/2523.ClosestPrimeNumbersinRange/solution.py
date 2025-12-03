from typing import List


class Solution:
    def closestPrimes(self, left: int, right: int) -> List[int]:
        if right < 2:
            return [-1, -1]

        is_comp = [False] * (right + 1)
        limit = int(right ** 0.5)
        for i in range(2, limit + 1):
            if not is_comp[i]:
                step_start = i * i
                for j in range(step_start, right + 1, i):
                    is_comp[j] = True

        prev = -1
        best_gap = 10**9
        ans = [-1, -1]

        for x in range(max(left, 2), right + 1):
            if not is_comp[x]:
                if prev != -1:
                    gap = x - prev
                    if gap < best_gap:
                        best_gap = gap
                        ans = [prev, x]
                prev = x

        return ans
