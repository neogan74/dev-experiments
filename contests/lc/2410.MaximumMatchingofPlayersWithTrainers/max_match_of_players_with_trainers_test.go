package maximummatchingofplayerswithtrainers

import "testing"

func Test_matchPlayersAndTrainers(t *testing.T) {
	type args struct {
		players  []int
		trainers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				players:  []int{4, 7, 9},
				trainers: []int{8, 2, 5, 8},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				players:  []int{1, 1, 1},
				trainers: []int{10},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchPlayersAndTrainers(tt.args.players, tt.args.trainers); got != tt.want {
				t.Errorf("matchPlayersAndTrainers() = %v, want %v", got, tt.want)
			}
		})
	}
}
