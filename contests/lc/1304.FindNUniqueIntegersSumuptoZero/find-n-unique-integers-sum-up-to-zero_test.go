package findnuniqueintegerssumuptozero

import (
	"reflect"
	"testing"
)

func Test_sumZero(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want []int
	}{
		{
			name: "Test 1",
			n:    5,
			want: []int{1, -1, 2, -2, 0},
		},
		{
			name: "Test 2",
			n:    3,
			want: []int{1, -1, 0},
		},
		{
			name: "Test 3",
			n:    1,
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumZero(tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				// TODO: update the condition below to compare got with tt.want.
				t.Errorf("sumZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
