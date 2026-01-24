package minimizemaximumpairsuminarray

import "testing"

type testCase struct {
	name  string
	input []int
	want  int
}

func TestMinPairSum(t *testing.T) {
	tests := []testCase{
		{
			name:  "example_one",
			input: []int{3, 5, 2, 3},
			want:  7,
		},
		{
			name:  "example_two",
			input: []int{3, 5, 4, 2, 4, 6},
			want:  8,
		},
		{
			name:  "sorted_pairs",
			input: []int{1, 2, 3, 4},
			want:  5,
		},
		{
			name:  "duplicate_values",
			input: []int{4, 4, 4, 4},
			want:  8,
		},
		{
			name:  "unbalanced_values",
			input: []int{1, 100000, 2, 99999},
			want:  100001,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			inputCopy := append([]int(nil), tc.input...)
			if got := minPairSum(inputCopy); got != tc.want {
				t.Fatalf("minPairSum(%v) = %d, want %d", tc.input, got, tc.want)
			}
		})
		t.Run(tc.name, func(t *testing.T) {
			inputCopy := append([]int(nil), tc.input...)
			if got := minPairSum2(inputCopy); got != tc.want {
				t.Fatalf("minPairSum2(%v) = %d, want %d", tc.input, got, tc.want)
			}
		})
	}
}
