package _1857_largest_color_graph

import "testing"

func Test_largestPathValue(t *testing.T) {
	tests := []struct {
		colors   string
		edges    [][]int
		expected int
	}{
		{
			colors:   "abaca",
			edges:    [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}},
			expected: 3,
		},
		{
			colors:   "a",
			edges:    [][]int{{0, 0}},
			expected: -1,
		},
		{
			colors:   "abc",
			edges:    [][]int{},
			expected: 1,
		},
		{
			colors:   "aaa",
			edges:    [][]int{{0, 1}, {1, 2}},
			expected: 3,
		},
	}

	for _, test := range tests {
		result := largestPathValue(test.colors, test.edges)
		if result != test.expected {
			t.Errorf("Expected %d, got %d for input colors=%s edges=%v", test.expected, result, test.colors, test.edges)
		}
	}
}
