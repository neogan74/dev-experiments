package _2094_finding_3_dig_even_num

import "sort"

func findNumbers(digits []int) []int {
	count := [10]int{}
	for _, d := range digits {
		count[d]++
	}

	var result []int
	for i := 100; i <= 998; i += 2 {
		a, b, c := i/100, (i/10)%10, i%10
		if count[a] > 0 && count[b] > 0 && count[c] > 0 {
			// временно уменьшаем счётчики, чтобы проверить уникальность
			count[a]--
			count[b]--
			count[c]--
			if count[a] >= 0 && count[b] >= 0 && count[c] >= 0 {
				result = append(result, i)
			}
			// восстанавливаем
			count[a]++
			count[b]++
			count[c]++
		}
	}

	sort.Ints(result)
	return result
}
