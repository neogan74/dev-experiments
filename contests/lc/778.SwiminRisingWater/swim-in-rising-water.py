from __future__ import annotations

from collections import deque

class Solution:
    def swimInWater(self, grid: list[list[int]]) -> int:
        n = len(grid)

        def can_reach(limit: int) -> bool:
            if grid[0][0] > limit:
                return False

            q = deque([(0, 0)])
            seen = [[False] * n for _ in range(n)]
            seen[0][0] = True

            while q:
                r, c = q.popleft()
                if (r, c) == (n - 1, n - 1):
                    return True
                for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                    nr, nc = r + dr, c + dc
                    if 0 <= nr < n and 0 <= nc < n and not seen[nr][nc] and grid[nr][nc] <= limit:
                        seen[nr][nc] = True
                        q.append((nr, nc))
            return False

        lo, hi = 0, n * n - 1
        while lo < hi:
            mid = (lo + hi) // 2
            if can_reach(mid):
                hi = mid
            else:
                lo = mid + 1
        return lo

if __name__ == "__main__":
    grid = [[0,2],[1,3]]
    sol = Solution()
    res = sol.swimInWater(grid=grid)
    print(f'Result for {grid} is {res}')
