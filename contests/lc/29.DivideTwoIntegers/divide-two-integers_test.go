package dividetwointegers

import "testing"

func Test_divide(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		a    int
		b    int
		want int
	}{
		{
			name: "Example 1",
			a:    10,
			b:    3,
			want: 3,
		},
		{
			name: "Example 2",
			a:    7,
			b:    -3,
			want: -2,
		},
		{
			name: "Example 3",
			a:    0,
			b:    1,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divide(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
