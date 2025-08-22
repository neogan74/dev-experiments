package findtheminimumareatocoverallonesii

import "testing"

func Test_minimumSum(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				grid: [][]int{{1, 0, 1}, {1, 1, 1}},
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				grid: [][]int{{1, 0, 1, 0}, {0, 1, 0, 1}},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSum(tt.args.grid); got != tt.want {
				t.Errorf("minimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
