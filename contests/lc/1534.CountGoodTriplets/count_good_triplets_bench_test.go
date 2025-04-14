package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateTestData(n, maxVal int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(maxVal)
	}
	return arr
}

func BenchmarkCountGoodTripletsNaive(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	arr := generateTestData(100, 100)
	a, bVal, c := 10, 10, 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countGoodTriplets(arr, a, bVal, c)
	}
}

func BenchmarkCountGoodTripletsOptimized(b *testing.B) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	arr := generateTestData(100, 100)
	a, bVal, c := 10, 10, 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		countGoodTripletsOptimized(arr, a, bVal, c)
	}
}
