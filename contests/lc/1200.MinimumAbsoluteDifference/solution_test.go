package minimumabsolutedifference

import "testing"

func Test_getMinimumAbsoluteDifference(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arr  []int
		want int
	}{
		{
			name: "Test 1",
			arr:  []int{4, 2, 1, 3},
			want: [][]int{{1, 2}, {2, 3}, {3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getMinimumAbsoluteDifference(tt.arr)
			if len(got) != len(tt.want) {
				t.Errorf("getMinimumAbsoluteDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
