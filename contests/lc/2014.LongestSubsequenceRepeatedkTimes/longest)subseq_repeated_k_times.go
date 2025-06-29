package longestsubsequencerepeatedktimes

func longestSubsequenceRepeatedK(s string, k int) string {
	// Подсчёт частот букв
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// Формируем кандидаты
	letters := []byte{}
	for c := byte('a'); c <= byte('z'); c++ {
		if freq[c] >= k {
			letters = append(letters, c)
		}
	}

	candidates := []string{""}
	maxLen := len(s) / k
	result := ""

	for length := 1; length <= maxLen; length++ {
		next := []string{}
		for _, t := range candidates {
			for _, c := range letters {
				newCandidate := t + string(c)
				rep := ""
				for i := 0; i < k; i++ {
					rep += newCandidate
				}
				if isSubsequence(s, rep) {
					next = append(next, newCandidate)
					if len(newCandidate) > len(result) ||
						(len(newCandidate) == len(result) && newCandidate > result) {
						result = newCandidate
					}
				}
			}
		}
		candidates = next
		if len(candidates) == 0 {
			break
		}
	}

	return result
}

func isSubsequence(haystack, needle string) bool {
	j := 0
	for i := 0; i < len(haystack) && j < len(needle); i++ {
		if haystack[i] == needle[j] {
			j++
		}
	}
	return j == len(needle)
}
