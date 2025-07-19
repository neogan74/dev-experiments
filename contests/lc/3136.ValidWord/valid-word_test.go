package validword

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{
				word: "234Adas",
			},
			want: true,
		},
		{
			name: "Test2",
			args: args{
				word: "b3",
			},
			want: false,
		},
		{
			name: "Test3",
			args: args{
				word: "a3$e",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.word); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
