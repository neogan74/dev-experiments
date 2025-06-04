package _2942_find_words_conitaining_character

func findWordsContaining(words []string, x byte) []int {
	var res []int
	for i, word := range words {
		for _, ch := range word {
			if ch == rune(x) {
				res = append(res, i)
				break
			}
		}
	}
	return res
}
