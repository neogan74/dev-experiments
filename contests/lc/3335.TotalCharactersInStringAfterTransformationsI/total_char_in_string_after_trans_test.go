package _3335_total_characters_in_string_after_trans

import "testing"

func Test_totalCharacters(t *testing.T) {
	type args struct {
		s string
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				s: "abcyy",
				t: 2,
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				s: "azbk",
				t: 1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalCharacters(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("totalCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
