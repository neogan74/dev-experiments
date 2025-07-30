package longestsubarraywithmaximumbitwiseand

import "testing"

func TestLongestSubarrayWithMaximumBitwiseAnd(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{arr: []int{1, 2, 3, 3, 2, 2}},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{arr: []int{1, 2, 3, 4}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestSubarrayWithMaximumBitwiseAnd(tt.args.arr); got != tt.want {
				t.Errorf("LongestSubarrayWithMaximumBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
