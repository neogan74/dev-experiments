from typing import List


class Solution:
    def maxMatrixSum(self, matrix: List[List[int]]) -> int:
        total = 0
        neg_count = 0
        min_abs = 10**18
        zero_found = False

        for row in matrix:
            for val in row:
                if val < 0:
                    neg_count += 1
                if val == 0:
                    zero_found = True
                abs_val = abs(val)
                total += abs_val
                if abs_val < min_abs:
                    min_abs = abs_val

        if neg_count % 2 == 1 and not zero_found:
            total -= 2 * min_abs
        return total
