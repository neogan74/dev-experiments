package _3341_find_minimum_time_to_reach_last_room_i

import "testing"

func Test_minTimeToReach(t *testing.T) {
	type args struct {
		moveTime [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				moveTime: [][]int{
					{0, 4},
					{4, 4},
				},
			},
			want: 6,
		},
		{
			name: "test2",
			args: args{
				moveTime: [][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToReach(tt.args.moveTime); got != tt.want {
				t.Errorf("minTimeToReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
