package _1298_max_candies

import "testing"

func Test_maxCandies(t *testing.T) {
	type args struct {
		status         []int
		candies        []int
		keys           [][]int
		containedBoxes [][]int
		initialBoxes   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				status:         []int{1, 0, 1, 0},
				candies:        []int{7, 5, 4, 100},
				keys:           [][]int{{}, {}, {1}, {}},
				containedBoxes: [][]int{{1, 2}, {3}, {}, {}},
				initialBoxes:   []int{0},
			},
			want: 16,
		},
		{
			name: "Test 2",
			args: args{
				status:         []int{1, 0, 0, 0, 0, 0},
				candies:        []int{1, 1, 1, 1, 1, 1},
				keys:           [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
				containedBoxes: [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}},
				initialBoxes:   []int{0},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCandies(tt.args.status, tt.args.candies, tt.args.keys, tt.args.containedBoxes, tt.args.initialBoxes); got != tt.want {
				t.Errorf("maxCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
