package _81_RabbitsinForest

import "testing"

func TestNumRabbits(t *testing.T) {
	tests := []struct {
		answers  []int
		expected int
	}{
		// Примеры и простые кейсы
		{[]int{}, 0},
		{[]int{0}, 1},
		{[]int{0, 0}, 2},
		{[]int{1, 1, 2}, 5},
		{[]int{0, 0, 1, 1, 1}, 6}, // две группы size=1 (от 0), две группы size=2 (от 1)

		// одна полная группа
		{[]int{2, 2, 2}, 3},    // groupSize=3, freq=3 → 1×3
		{[]int{3, 3, 3, 3}, 4}, // groupSize=4, freq=4 → 1×4

		// несколько групп
		{[]int{2, 2, 2, 2}, 6},          // groupSize=3, freq=4 → 2×3 = 6
		{[]int{2, 2, 2, 2, 2, 2, 2}, 9}, // freq=7, size=3 → 3×3 = 9

		// смешанные ответы
		{[]int{1, 0, 1, 0, 0}, 5}, // ans=1,freq=2→1group×2 + ans=0,freq=3→3×1 = 5
	}

	for _, tt := range tests {
		got := numRabbits(tt.answers)
		if got != tt.expected {
			t.Errorf("numRabbits(%v) = %d; want %d", tt.answers, got, tt.expected)
		}
	}
}
