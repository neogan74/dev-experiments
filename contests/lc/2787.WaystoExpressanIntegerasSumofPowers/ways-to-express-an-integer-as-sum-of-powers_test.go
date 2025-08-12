package waystoexpressanintegerassumofpowers

import "testing"

func Test_numberOfWays(t *testing.T) {
	type args struct {
		n int
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: 10,
				x: 2,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				n: 4,
				x: 1,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.args.n, tt.args.x); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
