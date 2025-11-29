package _512_MinimumOperationstoMakeArraySumDivisiblebyK

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "example 1",
			nums: []int{3, 9, 7},
			k:    5,
			want: 4,
		},
		{
			name: "example 2",
			nums: []int{4, 1, 3},
			k:    4,
			want: 0,
		},
		{
			name: "example 3",
			nums: []int{3, 2},
			k:    6,
			want: 5,
		},
		{
			name: "already divisible",
			nums: []int{2, 2, 2},
			k:    2,
			want: 0,
		},
		{
			name: "k equals 1",
			nums: []int{10, 20, 30},
			k:    1,
			want: 0,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.nums, tt.k)
			if got != tt.want {
				t.Fatalf("minOperations(%v, %d) = %d, want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
