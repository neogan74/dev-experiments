package besttimetobuyanssellstockusingstrategy

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		strategy []int
		k        int
		want     int64
	}{
		{
			name:     "example with modification",
			prices:   []int{4, 2, 8},
			strategy: []int{-1, 0, 1},
			k:        2,
			want:     10,
		},
		{
			name:     "example without modification",
			prices:   []int{5, 4, 3},
			strategy: []int{1, 1, 0},
			k:        2,
			want:     9,
		},
		{
			name:     "best window in the middle",
			prices:   []int{1, 2, 3, 4},
			strategy: []int{-1, -1, -1, -1},
			k:        2,
			want:     1,
		},
		{
			name:     "full window modification",
			prices:   []int{5, 1, 3, 4},
			strategy: []int{-1, 1, 0, 1},
			k:        4,
			want:     7,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices, tt.strategy, tt.k); got != tt.want {
				t.Fatalf("maxProfit() = %d, want %d", got, tt.want)
			}
		})
	}
}
