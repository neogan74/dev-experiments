package _2901_longest_uneq_groups_subsequence_ii

func getWordsInLongestSubsequence(words []string, groups []int) []string {
	n := len(words)

	// Функция: расстояние Хэмминга между словами
	hamming := func(a, b string) int {
		diff := 0
		for i := range a {
			if a[i] != b[i] {
				diff++
			}
		}
		return diff
	}

	// dp[i] — длина максимальной подпоследовательности, заканчивающейся на i
	dp := make([]int, n)
	prev := make([]int, n)

	for i := range dp {
		dp[i] = 1
		prev[i] = -1
	}

	// DP: перебираем все j < i
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if groups[i] != groups[j] &&
				len(words[i]) == len(words[j]) &&
				hamming(words[i], words[j]) == 1 {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					prev[i] = j
				}
			}
		}
	}

	// Восстановление самой длинной подпоследовательности
	bestLen, lastIdx := 0, -1
	for i, length := range dp {
		if length > bestLen {
			bestLen = length
			lastIdx = i
		}
	}

	var result []string
	for lastIdx != -1 {
		result = append([]string{words[lastIdx]}, result...)
		lastIdx = prev[lastIdx]
	}

	return result
}
