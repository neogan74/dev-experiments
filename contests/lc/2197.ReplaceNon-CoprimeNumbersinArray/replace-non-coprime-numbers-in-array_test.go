package replacenoncoprimenumbersinarray

import (
	"reflect"
	"testing"
)

func Test_replaceNonCoprimes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{6, 4, 3, 2, 7, 6, 2},
			want: []int{12, 7, 6},
		},
		{
			name: "Example 2",
			nums: []int{2, 2, 1, 1, 3, 3, 3},
			want: []int{2, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceNonCoprimes(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceNonCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
