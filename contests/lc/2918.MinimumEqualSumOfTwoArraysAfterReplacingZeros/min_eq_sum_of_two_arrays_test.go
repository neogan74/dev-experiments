package _2918_min_eq_sum_of_two_arrays_after_replacing_zeros

import "testing"

func Test_minSum(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Test 1",
			args: args{
				nums1: []int{3, 2, 0, 1, 0},
				nums2: []int{6, 5, 0},
			},
			want: 12,
		},
		{
			name: "Test 2",
			args: args{
				nums1: []int{2, 0, 2, 0},
				nums2: []int{1, 4},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSum(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("minSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
