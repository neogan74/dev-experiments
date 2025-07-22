package maximumerasurevalue

import "testing"

func Test_maximumUniqueSubarray(t *testing.T) {
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
				nums: []int{4, 2, 4, 5, 6},
			},
			wantAns: 17,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5},
			},
			wantAns: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maximumUniqueSubarray(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
