package _338_CountTheNumberOfIdealArrays

// ideal_arrays_test.go

import "testing"

func TestIdealArrays(t *testing.T) {
	tests := []struct {
		n, maxValue int
		expected    int
	}{
		{2, 5, 10},        // пример 1
		{5, 3, 11},        // пример 2
		{1, 10000, 10000}, // длина 1: любое число
		{3, 3, 7},         // все идеальные массивы длины 3, значения ≤3
		{10000, 1, 1},     // все элементы =1, единственный массив
		{10, 100, 105568}, // стресс-проверка
	}

	for _, tt := range tests {
		got := idealArrays(tt.n, tt.maxValue)
		if got != tt.expected {
			t.Errorf("idealArrays(%d, %d) = %d; want %d",
				tt.n, tt.maxValue, got, tt.expected)
		}
	}
}
