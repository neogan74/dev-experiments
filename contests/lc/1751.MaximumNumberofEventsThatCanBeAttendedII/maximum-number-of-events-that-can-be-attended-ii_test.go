package maximumnumberofeventsthatcanbeattendedii

import "testing"

func Test_maxValue(t *testing.T) {
	type args struct {
		events [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 1}, {4, 5, 2}},
				k:      2,
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				events: [][]int{{1, 2, 4}, {3, 4, 3}, {2, 3, 10}},
				k:      2,
			},
			want: 10,
		},
		{
			name: "Example 3",
			args: args{
				events: [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}},
				k:      3,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.events, tt.args.k); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
