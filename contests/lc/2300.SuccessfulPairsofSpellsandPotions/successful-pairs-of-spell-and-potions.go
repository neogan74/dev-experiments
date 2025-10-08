package main

import "sort"

// successfulPairs computes how many potions can pair with each spell to reach the success threshold.
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	n := len(spells)
	m := len(potions)
	result := make([]int, n)

	for i, spell := range spells {
		if spell == 0 {
			// A zero-strength spell can only succeed if success is also zero (per constraints success >= 1).
			result[i] = 0
			continue
		}

		// Minimum potion strength needed so spell * potion >= success.
		required := int64(spell)
		threshold := (success + required - 1) / required

		idx := sort.Search(m, func(j int) bool {
			return potions[j] >= int(threshold)
		})

		result[i] = m - idx
	}

	return result
}
