package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s := "the sky is blue"
	res := reverseWords(s)
	fmt.Println(res)
	s2 := "  hello world  "
	res = reverseWords(s2)
	fmt.Println(res)
	s3 := "a good   example"
	res = reverseWords(s3)
	fmt.Println(res)

}
