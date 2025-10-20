package _625_LexicographicallySmallestStringAfterApplyingOperations

import "testing"

func Test_findLexSmallestString(t *testing.T) {
	type args struct {
		s string
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				s: "5525",
				a: 9,
				b: 2,
			},
			want: "2050",
		},
		{
			name: "Test 2",
			args: args{
				s: "74",
				a: 5,
				b: 1,
			},
			want: "24",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLexSmallestString(tt.args.s, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("findLexSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
