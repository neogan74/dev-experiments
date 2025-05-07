package _2302_count_subarr_with_score

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
		k    int64
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{2, 1, 4, 3, 5},
				k:    10,
			},
			wantAns: 6,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 1, 1},
				k:    5,
			},
			wantAns: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSubarrays(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("countSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
