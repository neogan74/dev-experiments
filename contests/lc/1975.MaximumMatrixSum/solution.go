package main

func maxMatrixSum(matrix [][]int) int64 {
	var sum int64
	negCount := 0
	minAbs := int64(1<<63 - 1)
	zeroFound := false

	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		for j := 0; j < len(row); j++ {
			val := row[j]
			if val < 0 {
				negCount++
			}
			if val == 0 {
				zeroFound = true
			}
			absVal := int64(val)
			if absVal < 0 {
				absVal = -absVal
			}
			sum += absVal
			if absVal < minAbs {
				minAbs = absVal
			}
		}
	}

	if negCount%2 == 1 && !zeroFound {
		sum -= 2 * minAbs
	}
	return sum
}
