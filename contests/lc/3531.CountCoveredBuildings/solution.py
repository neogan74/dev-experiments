from typing import List


class Solution:
    def countCoveredBuildings(self, n: int, buildings: List[List[int]]) -> int:
        # Track the smallest and largest column per row, and row per column.
        row_limits: dict[int, tuple[int, int]] = {}
        col_limits: dict[int, tuple[int, int]] = {}

        for r, c in buildings:
            if r in row_limits:
                row_min, row_max = row_limits[r]
                row_limits[r] = (min(row_min, c), max(row_max, c))
            else:
                row_limits[r] = (c, c)

            if c in col_limits:
                col_min, col_max = col_limits[c]
                col_limits[c] = (min(col_min, r), max(col_max, r))
            else:
                col_limits[c] = (r, r)

        covered = 0
        for r, c in buildings:
            row_min, row_max = row_limits[r]
            col_min, col_max = col_limits[c]

            if c != row_min and c != row_max and r != col_min and r != col_max:
                covered += 1

        return covered
