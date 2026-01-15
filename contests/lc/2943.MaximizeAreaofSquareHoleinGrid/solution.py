from typing import List


class Solution:
    def maximizeSquareHoleArea(
        self, n: int, m: int, hBars: List[int], vBars: List[int]
    ) -> int:
        max_h = self._max_consecutive_span(hBars)
        max_v = self._max_consecutive_span(vBars)
        side = min(max_h, max_v)
        return side * side

    @staticmethod
    def _max_consecutive_span(bars: List[int]) -> int:
        if not bars:
            return 1
        bars.sort()
        longest = 1
        current = 1
        for i in range(1, len(bars)):
            if bars[i] == bars[i - 1] + 1:
                current += 1
            else:
                if current > longest:
                    longest = current
                current = 1
        if current > longest:
            longest = current
        return longest + 1
