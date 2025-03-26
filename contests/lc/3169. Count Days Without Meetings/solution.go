package main

import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
	// сортировка встреч по началу интервалов
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	totalBusy := 0
	currStart, currEnd := meetings[0][0], meetings[0][1]

	for _, meet := range meetings[1:] {
		start, end := meet[0], meet[1]
		if start <= currEnd {
			// пересекающиеся интервалы, расширяем текущий
			if end > currEnd {
				currEnd = end
			}
		} else {
			// нет пересечения, добавляем текущий интервал
			totalBusy += currEnd - currStart + 1
			currStart, currEnd = start, end
		}
	}
	// добавляем последний интервал
	totalBusy += currEnd - currStart + 1

	return days - totalBusy
}

func main() {
	fmt.Println(countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}})) // 2
	fmt.Println(countDays(5, [][]int{{2, 4}, {1, 3}}))           // 1
	fmt.Println(countDays(6, [][]int{{1, 6}}))                   // 0
}
