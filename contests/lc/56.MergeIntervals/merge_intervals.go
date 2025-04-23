package _6_MergeIntervals

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// Сортируем по начальному значению интервала
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for _, curr := range intervals[1:] {
		last := merged[len(merged)-1]
		if curr[0] <= last[1] {
			// пересекается — объединяем
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			// не пересекается — добавляем
			merged = append(merged, curr)
		}
	}

	return merged
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}
