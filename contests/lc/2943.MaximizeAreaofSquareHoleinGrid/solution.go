package main

import "sort"

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	maxH := maxConsecutiveSpan(hBars)
	maxV := maxConsecutiveSpan(vBars)
	if maxH < maxV {
		return maxH * maxH
	}
	return maxV * maxV
}

func maxConsecutiveSpan(bars []int) int {
	if len(bars) == 0 {
		return 1
	}
	sort.Ints(bars)
	longest := 1
	current := 1
	for i := 1; i < len(bars); i++ {
		if bars[i] == bars[i-1]+1 {
			current++
		} else {
			if current > longest {
				longest = current
			}
			current = 1
		}
	}
	if current > longest {
		longest = current
	}
	return longest + 1
}
