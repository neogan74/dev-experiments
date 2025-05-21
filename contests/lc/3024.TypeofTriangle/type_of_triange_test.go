package _3024_type_of_triange

import "testing"

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{3, 3, 3},
			},
			want: "equilateral",
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{3, 4, 5},
			},
			want: "scalene",
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1, 1, 1},
			},
			want: "equilateral",
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: "none",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := triangleType(tt.args.nums); got != tt.want {
				t.Errorf("triangleType() = %v, want %v", got, tt.want)
			}
		})
	}
}
