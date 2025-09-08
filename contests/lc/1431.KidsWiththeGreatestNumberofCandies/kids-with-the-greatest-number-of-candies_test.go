package kidswiththegreatestnumberofcandies

import (
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies     []int
		extra       int
		expectedAns []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},    // Стандартный случай
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}}, // Только первый ребенок может стать самым богатым
		{[]int{1, 1, 1, 1}, 0, []bool{true, true, true, true}},              // Все равны, extra = 0
		{[]int{5, 5, 5, 5}, 5, []bool{true, true, true, true}},              // Все получают достаточно, чтобы быть максимальными
		{[]int{0}, 0, []bool{true}},                                         // Один ребенок
		{[]int{0}, 1, []bool{true}},                                         // Один ребенок с extra                                          // Пустой массив (необычно, но допустимо)
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := kidsWithCandies(test.candies, test.extra)
			if len(result) != len(test.expectedAns) {
				t.Errorf("kidsWithCandies(%v, %d) = %v, want %v", test.candies, test.extra, result, test.expectedAns)
				return
			}
			for i := range result {
				if result[i] != test.expectedAns[i] {
					t.Errorf("kidsWithCandies(%v, %d) = %v, want %v", test.candies, test.extra, result, test.expectedAns)
					return
				}
			}
		})
	}
}
