package main

import (
	"fmt"
)

func findSubstring(s string, words []string) []int {
	result := []int{}
	if len(words) == 0 || len(s) == 0 {
		return result
	}

	wordLen := len(words[0])
	numWords := len(words)
	totalLen := wordLen * numWords
	if len(s) < totalLen {
		return result
	}

	// Построим мапу частот для слов
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	// Обходим возможные смещения от 0 до wordLen-1
	for i := 0; i < wordLen; i++ {
		// left – начало окна, right – конец текущего слова (в блоках)
		left := i
		currentCount := make(map[string]int)
		count := 0 // количество слов в текущем окне, которые "подсчитаны"

		// двигаем right, увеличивая окно на слово за раз
		for right := i; right+wordLen <= len(s); right += wordLen {
			word := s[right : right+wordLen]
			// Если слово присутствует в словаре
			if _, ok := wordCount[word]; ok {
				currentCount[word]++
				count++
				// Если слово встречается слишком часто, сдвигаем окно с левой стороны
				for currentCount[word] > wordCount[word] {
					// удаляем слово, которое вышло из окна
					leftWord := s[left : left+wordLen]
					currentCount[leftWord]--
					left += wordLen
					count--
				}
				// Если окно содержит ровно numWords слов, проверяем ответ
				if count == numWords {
					result = append(result, left)
					// Сдвигаем окно на один блок слова, чтобы искать следующее совпадение
					leftWord := s[left : left+wordLen]
					currentCount[leftWord]--
					left += wordLen
					count--
				}
			} else {
				// Если слово не из списка, сбрасываем окно
				currentCount = make(map[string]int)
				count = 0
				left = right + wordLen
			}
		}
	}

	return result
}

func main() {
	// Пример 1:
	s1 := "barfoothefoobarman"
	words1 := []string{"foo", "bar"}
	// Возможные индексы: [0,9]
	fmt.Printf("s = %q, words = %v → %v\n", s1, words1, findSubstring(s1, words1))

	// Пример 2:
	s2 := "wordgoodgoodgoodbestword"
	words2 := []string{"word", "good", "best", "word"}
	// Ожидается: [] (нет валидного конкатенированного подстроки)
	fmt.Printf("s = %q, words = %v → %v\n", s2, words2, findSubstring(s2, words2))

	// Пример 3:
	s3 := "barfoofoobarthefoobarman"
	words3 := []string{"bar", "foo", "the"}
	// Возможные индексы: [6,9,12] или иные, зависящие от расположения слов
	fmt.Printf("s = %q, words = %v → %v\n", s3, words3, findSubstring(s3, words3))
}
