package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[string][]string)

	for _, str := range strs {
		key := sortString(str)
		anagramGroups[key] = append(anagramGroups[key], str)
	}

	result := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		result = append(result, group)
	}
	return result
}

// Вспомогательная функция для сортировки строки
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
