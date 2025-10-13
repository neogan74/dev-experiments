package _273_FindResultantArrayAfterRemovingAnagrams

import "sort"

func removeAnagrams(words []string) []string {
	result := make([]string, 0, len(words))
	var prevKey string
	hasPrev := false

	for _, w := range words {
		key := sortString(w)
		if !hasPrev || key != prevKey {
			result = append(result, w)
			prevKey = key
			hasPrev = true
		}
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
