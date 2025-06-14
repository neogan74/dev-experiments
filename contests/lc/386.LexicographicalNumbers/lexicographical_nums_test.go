package _386_lexiogrpahical_numbers

import (
	"reflect"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "Test 1",
			args: args{
				n: 13,
			},
			wantAns: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test 2",
			args: args{
				n: 2,
			},
			wantAns: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := lexicalOrder(tt.args.n); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("lexicalOrder() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
