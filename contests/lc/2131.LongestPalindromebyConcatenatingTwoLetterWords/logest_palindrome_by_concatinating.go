package _2131_longest_palindrome_by_concatinating_two_letter_words

func longestPalindrome(words []string) int {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	ans, x := 0, 0
	for k, v := range cnt {
		if k[0] == k[1] {
			x += v & 1
			ans += v / 2 * 2 * 2
		} else {
			if y, ok := cnt[string(k[1])+string(k[0])]; ok {
				ans += min(v, y) * 2
			}
		}
	}
	if x > 0 {
		ans += 2
	}
	return ans
}
