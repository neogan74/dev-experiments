package countnumberoftrapezoidsii

import "testing"

func Test_numberOfTrapezoids(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		points [][]int
		want   int
	}{
		{
			name: "example 1",
			points: [][]int{
				{0, 0},
				{2, 0},
				{1, 1},
				{3, 1},
			},
			want: 1,
		},
		{
			name: "example 2",
			points: [][]int{
				{-3, 2},
				{3, 0},
				{2, 3},
				{3, 2},
				{2, -3},
			},
			want: 2,
		},
		{
			name: "example 3",
			points: [][]int{
				{0, 0},
				{1, 0},
				{0, 1},
				{2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfTrapezoids(tt.points)
			if got != tt.want {
				t.Errorf("numberOfTrapezoids() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			got := numberOfTrapezoids2(tt.points)
			if got != tt.want {
				t.Errorf("numberOfTrapezoids() = %v, want %v", got, tt.want)
			}
		})
	}
}
