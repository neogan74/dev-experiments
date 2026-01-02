package repeatedElementinSize2NArray

import "testing"

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example-1",
			nums: []int{1, 2, 3, 3},
			want: 3,
		},
		{
			name: "example-2",
			nums: []int{2, 1, 2, 5, 3, 2},
			want: 2,
		},
		{
			name: "example-3",
			nums: []int{5, 1, 5, 2, 5, 3, 5, 4},
			want: 5,
		},
		{
			name: "repeat-at-start",
			nums: []int{9, 9, 1, 2},
			want: 9,
		},
		{
			name: "repeat-at-end",
			nums: []int{4, 1, 2, 4},
			want: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := repeatedNTimes(test.nums); got != test.want {
				t.Fatalf("repeatedNTimes(%v) = %d, want %d", test.nums, got, test.want)
			}
		})
	}
}
