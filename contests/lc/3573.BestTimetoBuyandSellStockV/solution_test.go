package _besttimeofbiyandsellstockv

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		k      int
		prices []int
		want   int
	}{
		{
			name:   "example1",
			k:      2,
			prices: []int{1, 7, 9, 8, 2},
			want:   14,
		},
		{
			name:   "example2",
			k:      3,
			prices: []int{12, 16, 19, 19, 8, 1, 19, 13, 9},
			want:   36,
		},
		{
			name:   "no transactions allowed",
			k:      0,
			prices: []int{5, 4, 3},
			want:   0,
		},
		{
			name:   "single price",
			k:      1,
			prices: []int{5},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProfit(tt.prices, tt.k); got != tt.want {
				t.Fatalf("maximumProfit(%v, %d) = %d, want %d", tt.prices, tt.k, got, tt.want)
			}
		})
	}
}
