package main

import (
	"math/rand"
	"testing"
	"time"
)

// Создание анаграммы заданной длины из случайных букв [a-z]
func generateAnagramPair(length int) (string, string) {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	s := make([]rune, length)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	t := append([]rune{}, s...)
	rand.Shuffle(len(t), func(i, j int) { t[i], t[j] = t[j], t[i] })
	return string(s), string(t)
}

func BenchmarkIsAnagramArray(b *testing.B) {
	s, t := generateAnagramPair(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagram(s, t)
	}
}

func BenchmarkIsAnagramMap(b *testing.B) {
	s, t := generateAnagramPair(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isAnagramUnicode(s, t)
	}
}
