from typing import List


class Solution:
    def largestSquareArea(self, bottomLeft: List[List[int]], topRight: List[List[int]]) -> int:
        max_side = 0
        n = len(bottomLeft)
        for i in range(n):
            for j in range(i + 1, n):
                left = max(bottomLeft[i][0], bottomLeft[j][0])
                right = min(topRight[i][0], topRight[j][0])
                bottom = max(bottomLeft[i][1], bottomLeft[j][1])
                top = min(topRight[i][1], topRight[j][1])

                width = right - left
                height = top - bottom
                if width <= 0 or height <= 0:
                    continue

                side = min(width, height)
                if side > max_side:
                    max_side = side

        return max_side * max_side
