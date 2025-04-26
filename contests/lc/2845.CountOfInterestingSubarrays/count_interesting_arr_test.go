package _845_CountOfInterestingSubarrays

import "testing"

func Test_countInterestingSubarrays(t *testing.T) {
	type args struct {
		nums   []int
		modulo int
		k      int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name: "Example 1",
			args: args{
				nums:   []int{3, 2, 4},
				modulo: 2,
				k:      1,
			},
			wantAns: 3,
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{3, 1, 9, 6},
				modulo: 3,
				k:      0,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countInterestingSubarrays(tt.args.nums, tt.args.modulo, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("countInterestingSubarrays() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
