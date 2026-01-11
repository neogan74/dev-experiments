package main

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	cols := len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for _, row := range matrix {
		for j := 0; j < cols; j++ {
			if row[j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func largestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([]int, 0, len(heights))

	for i := 0; i <= len(heights); i++ {
		curr := 0
		if i < len(heights) {
			curr = heights[i]
		}

		for len(stack) > 0 && curr < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			height := heights[top]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}
