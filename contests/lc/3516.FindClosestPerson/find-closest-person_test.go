package findclosestperson

import "testing"

func Test_findClosest(t *testing.T) {
	type args struct {
		x int
		y int
		z int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				x: 2,
				y: 7,
				z: 4,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				x: 2,
				y: 5,
				z: 6,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				x: 1,
				y: 5,
				z: 3,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findClosest(tt.args.x, tt.args.y, tt.args.z); got != tt.want {
				t.Errorf("findClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
