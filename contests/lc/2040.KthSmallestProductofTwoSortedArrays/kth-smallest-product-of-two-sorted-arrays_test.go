package _2040_kth_smallest_product_of_two_sorted_arrays

import "testing"

func Test_kthSmallestProduct(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{2, 5},
				nums2: []int{3, 4},
				k:     2,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestProduct(tt.args.nums1, tt.args.nums2, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
