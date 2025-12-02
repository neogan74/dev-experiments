package numberoflaserbeamsinabank

import (
	"strings"
)

func countOnes(s string) int {
	return strings.Count(s, "1")
}

func numberOfBeams(bank []string) int {
	totalBeams := 0
	prevCount := 0

	for _, row := range bank {
		currentCount := countOnes(row)
		if currentCount == 0 {
			continue
		}
		totalBeams += prevCount * currentCount
		prevCount = currentCount
	}

	return totalBeams
}
