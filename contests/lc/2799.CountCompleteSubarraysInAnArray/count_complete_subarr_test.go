package _799_CountCompleteSubarraysInAnArray

import "testing"

func Test_countCompleteSubarrays(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{
				nums: []int{1, 3, 1, 2, 2},
			},
			wantAns: 4,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{5, 5, 5, 5},
			},
			wantAns: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countCompleteSubarrays(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countCompleteSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
