package _003_MaximizetheNumberofPartitionsAfterOperations

import "testing"

func Test_maxPartitionsAfterOperations(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "accca",
				k: 2,
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				s: "aabaab",
				k: 3,
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				s: "xxyz",
				k: 1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPartitionsAfterOperations(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxPartitionsAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
