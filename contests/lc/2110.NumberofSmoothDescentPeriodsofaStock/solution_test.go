package numberofsmoothdescentperiodsofastock

import "testing"

func TestCountDescentPeriods(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int64
	}{
		{name: "example 1", prices: []int{3, 2, 1, 4}, want: 7},
		{name: "example 2", prices: []int{8, 6, 7, 7}, want: 4},
		{name: "example 3", prices: []int{1}, want: 1},
		{name: "all descending", prices: []int{5, 4, 3, 2, 1}, want: 15}, // 5*6/2
		{name: "no descents", prices: []int{5, 5, 5}, want: 3},
		{name: "alternating", prices: []int{4, 3, 4, 3, 2}, want: 9},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := getDescentPeriods(tc.prices); got != tc.want {
				t.Fatalf("getDescentPeriods(%v) = %d, want %d", tc.prices, got, tc.want)
			}
		})
	}
}
