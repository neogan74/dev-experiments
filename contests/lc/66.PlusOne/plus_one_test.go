package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "Test 2",
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "Test 3",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := plusOne(test.args.digits)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("plusOne() = %v; want %v", got, test.want)
			}
		})
	}
}
