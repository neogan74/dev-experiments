package main

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
