package main

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	result := make([]string, n)

	// сохраняем пары (score, index)
	type pair struct {
		score int
		index int
	}
	pairs := make([]pair, n)
	for i := 0; i < n; i++ {
		pairs[i] = pair{score[i], i}
	}

	// сортируем по убыванию score
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].score > pairs[j].score
	})

	for rank, p := range pairs {
		switch rank {
		case 0:
			result[p.index] = "Gold Medal"
		case 1:
			result[p.index] = "Silver Medal"
		case 2:
			result[p.index] = "Bronze Medal"
		default:
			result[p.index] = strconv.Itoa(rank + 1)
		}
	}

	return result
}
