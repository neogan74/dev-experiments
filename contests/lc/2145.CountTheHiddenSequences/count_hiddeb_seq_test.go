package _145_CountTheHiddenSequences

import "testing"

func TestCountHiddenSequences(t *testing.T) {
	tests := []struct {
		differences  []int
		lower, upper int
		expected     int
	}{
		// примеры из условия
		{[]int{1, -3, 4}, 1, 6, 2},
		{[]int{3, -4, 5, 1, -2}, -4, 5, 4},
		{[]int{4, -7, 2}, 3, 6, 0},

		// все перепады нулевые → любые константные seq ∈ [lower, upper]
		{[]int{0, 0, 0}, 5, 5, 1}, // только [5,5,5,5]
		{[]int{0, 0}, 1, 3, 3},    // [1,1,1], [2,2,2], [3,3,3]

		// отрицательные и положительные перепады
		{[]int{-1, 2, -1}, 0, 2, 1},
		// prefix:0,-1,1,0 → minPref=-1,maxPref=1
		// x∈[0-(-1)=1, 2-1=1] → только x=1 → одна последовательность
		// но длина diff=3→len=4, countHiddenSequences вернёт (2-1)-(0-(-1))+1 =1-1+1=1

		// большие диапазоны
		{[]int{100000, 100000, -200000}, -100000, 200000, 100001},
		// prefix:0,100k,200k,0 → minPref=0,maxPref=200k
		// x∈[-100k-0=−100k, 200k-200k=0] → от -100k до 0 включительно → 100001 вариантов
	}

	for _, tt := range tests {
		got := countHiddenSequences(tt.differences, tt.lower, tt.upper)
		if got != tt.expected {
			t.Errorf("countHiddenSequences(%v, %d, %d) = %d; want %d",
				tt.differences, tt.lower, tt.upper, got, tt.expected)
		}
	}
}
