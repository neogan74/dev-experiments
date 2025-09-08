package bitwiseorsofsubarrays

import "testing"

func Test_subarrayBitwiseORs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				arr: []int{0},
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				arr: []int{1, 1, 2},
			},
			want: 3,
		},
		{
			name: "test3",
			args: args{
				arr: []int{1, 2, 4},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarrayBitwiseORs(tt.args.arr); got != tt.want {
				t.Errorf("subarrayBitwiseORs() = %v, want %v", got, tt.want)
			}
		})
	}
}
