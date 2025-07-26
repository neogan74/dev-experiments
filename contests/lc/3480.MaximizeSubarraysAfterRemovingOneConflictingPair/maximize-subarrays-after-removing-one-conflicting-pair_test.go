package maximizesubarraysafterremovingoneconflictingpair

import "testing"

func Test_maxSubarrays(t *testing.T) {
	type args struct {
		n                int
		conflictingPairs [][]int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "test1",
			args: args{
				n:                4,
				conflictingPairs: [][]int{{2, 3}, {1, 4}},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarrays(tt.args.n, tt.args.conflictingPairs); got != tt.want {
				t.Errorf("maxSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
