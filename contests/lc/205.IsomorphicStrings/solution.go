package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS2T := make(map[byte]byte)
	mapT2S := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		c1, c2 := s[i], t[i]

		if v, ok := mapS2T[c1]; ok {
			if v != c2 {
				return false
			}
		} else {
			mapS2T[c1] = c2
		}

		if v, ok := mapT2S[c2]; ok {
			if v != c1 {
				return false
			}
		} else {
			mapT2S[c2] = c1
		}
	}

	return true
}

func main() {
	tests := []struct {
		s, t string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
		{"badc", "baba", false},
	}

	for _, test := range tests {
		got := isIsomorphic(test.s, test.t)
		fmt.Printf("isIsomorphic(%q, %q) = %v (want %v)\n", test.s, test.t, got, test.want)
	}
}
