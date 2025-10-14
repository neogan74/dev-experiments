package _349_AdjacentIncreasingSubarraysDetectionI

import "testing"

func Test_hasUniqueElements(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasUniqueElements(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("hasUniqueElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
