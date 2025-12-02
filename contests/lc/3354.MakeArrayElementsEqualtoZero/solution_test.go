package _354_MakeArrayElementsEqualtoZero

import "testing"

func Test_countValidSelections(t *testing.T) {
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
				nums: []int{1, 0, 2, 0, 3},
			},
			wantAns: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{2, 3, 4, 0, 4, 1, 0},
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countValidSelections(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("countValidSelections() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
