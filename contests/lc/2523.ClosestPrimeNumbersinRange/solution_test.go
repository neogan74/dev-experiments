package closestprimenumbersinrange

import "testing"

func TestClosestPrimes(t *testing.T) {
	tests := []struct {
		name        string
		left, right int
		want        []int
	}{
		{
			name:  "sample gap two",
			left:  10,
			right: 19,
			want:  []int{11, 13},
		},
		{
			name:  "not enough primes",
			left:  4,
			right: 6,
			want:  []int{-1, -1},
		},
		{
			name:  "single prime only",
			left:  1,
			right: 2,
			want:  []int{-1, -1},
		},
		{
			name:  "tightest gap one",
			left:  2,
			right: 3,
			want:  []int{2, 3},
		},
		{
			name:  "choose earliest pair when gaps tie",
			left:  2,
			right: 11,
			want:  []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := closestPrimes(tt.left, tt.right)
			if len(got) != len(tt.want) {
				t.Fatalf("length mismatch: got %v want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Fatalf("closestPrimes(%d, %d) = %v, want %v", tt.left, tt.right, got, tt.want)
				}
			}
		})
	}
}
