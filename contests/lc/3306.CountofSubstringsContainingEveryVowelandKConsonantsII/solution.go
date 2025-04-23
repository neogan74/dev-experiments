package main

import (
	"fmt"
	"sort"
)

func numberOfSubstrings(word string, k int) int64 {
	n := len(word)
	// C[i] will store the number of consonants in word[0...i-1].
	C := make([]int, n+1)
	for i := 0; i < n; i++ {
		C[i+1] = C[i]
		if !isVowel(word[i]) {
			C[i+1]++
		}
	}

	// Build a map: for each possible consonant cumulative value, store the sorted list of indices l
	indicesByC := make(map[int][]int)
	for l := 0; l <= n; l++ {
		indicesByC[C[l]] = append(indicesByC[C[l]], l)
	}

	// lastOccurrence will record the last seen index for each vowel.
	lastOccurrence := map[byte]int{'a': -1, 'e': -1, 'i': -1, 'o': -1, 'u': -1}
	result := int64(0)

	// Iterate over r (end index of substring)
	for r := 0; r < n; r++ {
		ch := word[r]
		if isVowel(ch) {
			lastOccurrence[ch] = r
		}
		// Check if all vowels have been seen so far.
		if lastOccurrence['a'] == -1 || lastOccurrence['e'] == -1 ||
			lastOccurrence['i'] == -1 || lastOccurrence['o'] == -1 || lastOccurrence['u'] == -1 {
			continue
		}
		// The valid l must be <= Lmin.
		Lmin := lastOccurrence['a']
		for _, v := range []byte{'e', 'i', 'o', 'u'} {
			if lastOccurrence[v] < Lmin {
				Lmin = lastOccurrence[v]
			}
		}
		// target cumulative value for l: we need C[l] = C[r+1] - k.
		target := C[r+1] - k
		if target < 0 {
			continue // cannot have negative cumulative consonants.
		}
		// Look up indices for this target value.
		indices, exists := indicesByC[target]
		if !exists {
			continue
		}
		// Count how many indices in indices are <= Lmin.
		// Since indices is sorted, we can use binary search.
		count := sort.Search(len(indices), func(i int) bool {
			return indices[i] > Lmin
		})
		// count is the number of indices with indices[i] <= Lmin.
		result += int64(count)
	}

	return result
}

func countOfSubstrings2(word string, k int) int64 {
	isVowel := make([]bool, 128) // To mark vowels
	freq := make([]int, 128)     // To track character frequency
	vowels := "aeiou"

	// Mark vowels in the isVowel array
	for _, v := range vowels {
		isVowel[v] = true
	}

	var response int64 = 0
	currentK, vowelCount, extraLeft, left := 0, 0, 0, 0

	for right := 0; right < len(word); right++ {
		rightChar := word[right]

		if isVowel[rightChar] {
			freq[rightChar]++
			if freq[rightChar] == 1 {
				vowelCount++
			}
		} else {
			currentK++
		}

		// Shrink the window if consonant count exceeds k
		for currentK > k {
			leftChar := word[left]
			if isVowel[leftChar] {
				freq[leftChar]--
				if freq[leftChar] == 0 {
					vowelCount--
				}
			} else {
				currentK--
			}
			left++
			extraLeft = 0
		}

		// Adjust left pointer to remove extra vowels
		for vowelCount == 5 && currentK == k && left < right && isVowel[word[left]] && freq[word[left]] > 1 {
			extraLeft++
			freq[word[left]]--
			left++
		}

		// Count valid substrings
		if currentK == k && vowelCount == 5 {
			response += int64(1 + extraLeft)
		}
	}

	return response
}

// Helper function: returns true if ch is a vowel.
func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}

func main() {
	// Example 1
	word1 := "aeioqq"
	k1 := 1
	fmt.Println(numberOfSubstrings(word1, k1)) // Expected output: 0
	fmt.Println(countOfSubstrings2(word1, k1)) // Expected output: 0

	// Example 2
	word2 := "aeiou"
	k2 := 0
	fmt.Println(numberOfSubstrings(word2, k2)) // Expected output: 1
	fmt.Println(countOfSubstrings2(word2, k2)) // Expected output: 1

	// Example 3
	word3 := "ieaouqqieaouqq"
	k3 := 1
	fmt.Println(numberOfSubstrings(word3, k3)) // Expected output: 3
	fmt.Println(countOfSubstrings2(word3, k3)) // Expected output: 3
}
