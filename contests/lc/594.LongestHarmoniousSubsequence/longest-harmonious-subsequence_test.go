package longestharmonioussubsequence

import "testing"

func Test_findLHS(t *testing.T) {
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
				nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			},
			wantAns: 5,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			wantAns: 2,
		},
		{
			name: "Test 3",
			args: args{
				nums: []int{1, 1, 1, 1},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findLHS(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findLHS() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
