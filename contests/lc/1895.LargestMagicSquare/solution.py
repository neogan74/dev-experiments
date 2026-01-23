from typing import List


class Solution:
    def largestMagicSquare(self, grid: List[List[int]]) -> int:
        m = len(grid)
        if m == 0:
            return 0
        n = len(grid[0])

        row_pref = [[0] * (n + 1) for _ in range(m)]
        for i in range(m):
            for j in range(n):
                row_pref[i][j + 1] = row_pref[i][j] + grid[i][j]

        col_pref = [[0] * n for _ in range(m + 1)]
        for j in range(n):
            for i in range(m):
                col_pref[i + 1][j] = col_pref[i][j] + grid[i][j]

        diag_pref = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m):
            for j in range(n):
                diag_pref[i + 1][j + 1] = diag_pref[i][j] + grid[i][j]

        anti_pref = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m):
            for j in range(n):
                anti_pref[i + 1][j] = anti_pref[i][j + 1] + grid[i][j]

        max_size = min(m, n)
        for k in range(max_size, 1, -1):
            for i in range(m - k + 1):
                for j in range(n - k + 1):
                    target = row_pref[i][j + k] - row_pref[i][j]
                    diag = diag_pref[i + k][j + k] - diag_pref[i][j]
                    if diag != target:
                        continue
                    anti = anti_pref[i + k][j] - anti_pref[i][j + k]
                    if anti != target:
                        continue
                    ok = True
                    for r in range(k):
                        row_sum = row_pref[i + r][j + k] - row_pref[i + r][j]
                        if row_sum != target:
                            ok = False
                            break
                        col_sum = col_pref[i + k][j + r] - col_pref[i][j + r]
                        if col_sum != target:
                            ok = False
                            break
                    if ok:
                        return k
        return 1
