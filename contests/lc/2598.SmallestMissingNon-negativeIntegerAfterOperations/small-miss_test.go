package _598_SmallestMissingNon_negativeIntegerAfterOperations

import "testing"

func Test_findSmallestInteger(t *testing.T) {
	type args struct {
		nums  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				nums:  []int{1, -10, 7, 13, 6, 8},
				value: 5,
			},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{
				nums:  []int{1, -10, 7, 13, 6, 8},
				value: 7,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSmallestInteger(tt.args.nums, tt.args.value); got != tt.want {
				t.Errorf("findSmallestInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}
