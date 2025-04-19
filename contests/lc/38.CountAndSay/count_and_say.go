package _8_CountAndSay

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	res := "1"
	// Генерируем от 2 до n
	for i := 2; i <= n; i++ {
		res = nextSequence(res)
	}
	return res
}

// nextSequence строит «описание» предыдущей строки
func nextSequence(s string) string {
	var next string
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			next += strconv.Itoa(count)
			next += string(s[i-1])
			count = 1
		}
	}
	// добавить последнюю группу
	next += strconv.Itoa(count)
	next += string(s[len(s)-1])
	return next
}

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("n = %d → %s\n", i, countAndSay(i))
	}
}
