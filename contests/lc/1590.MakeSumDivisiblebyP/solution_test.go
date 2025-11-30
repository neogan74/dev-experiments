package main

import "testing"

func TestMinSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		p    int
		want int
	}{
		{
			name: "example1",
			nums: []int{3, 1, 4, 2},
			p:    6,
			want: 1,
		},
		{
			name: "example2",
			nums: []int{6, 3, 5, 2},
			p:    9,
			want: 2,
		},
		{
			name: "already divisible",
			nums: []int{1, 2, 3},
			p:    3,
			want: 0,
		},
		{
			name: "impossible without removing all",
			nums: []int{1, 2, 3},
			p:    7,
			want: -1,
		},
		{
			name: "single removal works",
			nums: []int{5, 7, 10},
			p:    3,
			want: 1,
		},
		{
			name: "no removal needed with large values",
			nums: []int{1_000_000_000, 2_000_000_000},
			p:    3,
			want: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := minSubarray(tc.nums, tc.p); got != tc.want {
				t.Fatalf("minSubarray(%v, %d) = %d, want %d", tc.nums, tc.p, got, tc.want)
			}
		})
		t.Run(tc.name, func(t *testing.T) {
			if got := minSubarray2(tc.nums, tc.p); got != tc.want {
				t.Fatalf("minSubarray2(%v, %d) = %d, want %d", tc.nums, tc.p, got, tc.want)
			}
		})
	}
}

func BenchmarkMinSubarray(b *testing.B) {
	// Build a sizable input to exercise the O(n) scan.
	n := 100_000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i%10 + 1
	}
	p := 97

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = minSubarray(nums, p)
	}
}

func BenchmarkMinSubarray2(b *testing.B) {
	// Build a sizable input to exercise the O(n) scan.
	n := 100_000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i%10 + 1
	}
	p := 97

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = minSubarray2(nums, p)
	}
}
