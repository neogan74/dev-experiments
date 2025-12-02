package _526_MinimumNumberofIncrementsonSubarraystoFormaTargetArray

import "testing"

func Test_minNumberOperations(t *testing.T) {
	type args struct {
		target []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				target: []int{1, 2, 3, 2, 1},
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				target: []int{3, 1, 1, 2},
			},
			want: 4,
		},
		{
			name: "Test 3",
			args: args{
				target: []int{3, 1, 5, 4, 2},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minNumberOperations(tt.args.target); got != tt.want {
				t.Errorf("minNumberOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
