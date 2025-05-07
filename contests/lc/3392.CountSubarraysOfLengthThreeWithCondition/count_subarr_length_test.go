package _392_CountSubarraysOfLengthThreeWithCondition

import "testing"

func Test_countSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 1, 4, 1},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countSubarrays(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
