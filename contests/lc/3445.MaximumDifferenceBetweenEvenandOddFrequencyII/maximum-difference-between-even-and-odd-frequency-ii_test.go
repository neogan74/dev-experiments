package maximumdifferencebetweenevenandoddfrequencyii

import "testing"

func Test_maximumDifference(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test case 1",
			args: args{s: "12233", k: 4},
			want: -1,
		},
		{name: "Test case 2",
			args: args{s: "1122211", k: 3},
			want: 1,
		},
		{name: "Test case 3",
			args: args{s: "110", k: 3},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDifference(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maximumDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
