package minimumpenaltyforashop

import (
	"strings"
)

func bestClosingTime(customers string) int {
	penalty := strings.Count(customers, "Y")
	minPenalty := penalty
	result := 0

	for i := 0; i < len(customers); i++ {
		if customers[i] == 'Y' {
			penalty--
		} else {
			penalty++
		}

		if penalty < minPenalty {
			minPenalty = penalty
			result = i + 1
		}
	}

	return result
}
