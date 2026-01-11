from typing import List


class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        if not matrix or not matrix[0]:
            return 0

        cols = len(matrix[0])
        heights = [0] * cols
        max_area = 0

        for row in matrix:
            for j in range(cols):
                if row[j] == "1":
                    heights[j] += 1
                else:
                    heights[j] = 0
            max_area = max(max_area, self._largest_rectangle_area(heights))

        return max_area

    def _largest_rectangle_area(self, heights: List[int]) -> int:
        max_area = 0
        stack: List[int] = []

        for i in range(len(heights) + 1):
            curr = heights[i] if i < len(heights) else 0
            while stack and curr < heights[stack[-1]]:
                top = stack.pop()
                height = heights[top]
                width = i if not stack else i - stack[-1] - 1
                max_area = max(max_area, height * width)
            stack.append(i)

        return max_area
