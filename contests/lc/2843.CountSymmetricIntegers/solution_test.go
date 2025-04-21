package main

import "testing"

func TestCountSymmetricIntegers(t *testing.T) {
	tests := []struct {
		low, high int
		expected  int
	}{
		{1, 100, 9},         // 11,22,...,99
		{10, 99, 9},         // 11,22,...,99
		{100, 300, 0},       // 121,132,...,303
		{1, 9, 0},           // нет чётных по длине
		{1000, 1100, 2},     // 1001, 1010, 1100, 1111, 1120
		{123, 123, 0},       // 123 — длина 3 (нечётная)
		{1221, 1221, 1},     // 1+2 == 2+1
		{100001, 100999, 3}, // тест на 6-значные
		{1000, 1100, 2},     // 1001, 1010, 1100, 1111, 1120
	}

	for _, tt := range tests {
		result := countSymmetricIntegers(tt.low, tt.high)
		if result != tt.expected {
			t.Errorf("countSymmetricIntegers(%d, %d) = %d; want %d",
				tt.low, tt.high, result, tt.expected)
		}
	}
}
