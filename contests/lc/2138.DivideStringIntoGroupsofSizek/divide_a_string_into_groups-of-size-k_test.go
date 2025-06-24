package dividestringintogroupsofsizek

import (
	"reflect"
	"testing"
)

func Test_divideString(t *testing.T) {
	type args struct {
		s    string
		k    int
		fill byte
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			name: "Test 1",
			args: args{
				s:    "abcdefghi",
				k:    3,
				fill: byte('x'),
			},
			wantAns: []string{"abc", "def", "ghi"},
		},
		{
			name: "Test 1",
			args: args{
				s:    "abcdefghij",
				k:    3,
				fill: byte('x'),
			},
			wantAns: []string{"abc", "def", "ghi", "jxx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := divideString(tt.args.s, tt.args.k, tt.args.fill); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("divideString() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
