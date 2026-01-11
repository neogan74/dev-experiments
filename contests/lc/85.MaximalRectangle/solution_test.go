package main

import "testing"

func toMatrix(rows []string) [][]byte {
	matrix := make([][]byte, len(rows))
	for i, row := range rows {
		matrix[i] = []byte(row)
	}
	return matrix
}

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			name: "example1",
			matrix: toMatrix([]string{
				"10100",
				"10111",
				"11111",
				"10010",
			}),
			want: 6,
		},
		{
			name:   "example2",
			matrix: toMatrix([]string{"0"}),
			want:   0,
		},
		{
			name:   "example3",
			matrix: toMatrix([]string{"1"}),
			want:   1,
		},
		{
			name: "allZeros",
			matrix: toMatrix([]string{
				"000",
				"000",
			}),
			want: 0,
		},
		{
			name: "allOnes",
			matrix: toMatrix([]string{
				"11",
				"11",
				"11",
			}),
			want: 6,
		},
		{
			name:   "emptyMatrix",
			matrix: [][]byte{},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.matrix); got != tt.want {
				t.Fatalf("maximalRectangle() = %d, want %d", got, tt.want)
			}
		})
	}
}
