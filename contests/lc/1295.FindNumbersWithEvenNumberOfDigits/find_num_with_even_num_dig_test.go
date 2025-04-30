package __1295_FindNumbersWithEvenNumberOfDigits

import "testing"

func Test_findNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			"test1",
			args{[]int{12, 345, 2, 6, 7896}},
			2,
		},
		{
			"test2",
			args{[]int{555, 901, 482, 1771}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findNumbers(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
