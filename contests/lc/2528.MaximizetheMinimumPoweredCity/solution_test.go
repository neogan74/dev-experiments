package maximizetheminimumpoweredcity

import "testing"

func TestMaxPower(t *testing.T) {
	tests := []struct {
		name     string
		stations []int
		r        int
		k        int
		want     int64
	}{
		{
			name:     "example 1",
			stations: []int{1, 2, 4, 5, 0},
			r:        1,
			k:        2,
			want:     5,
		},
		{
			name:     "example 2",
			stations: []int{4, 4, 4, 4},
			r:        0,
			k:        3,
			want:     4,
		},
		{
			name:     "range spreads coverage",
			stations: []int{0, 0, 0},
			r:        1,
			k:        3,
			want:     3,
		},
		{
			name:     "large single city",
			stations: []int{100000},
			r:        0,
			k:        0,
			want:     100000,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.stations, tt.r, tt.k); got != tt.want {
				t.Fatalf("maxPower(%v, %d, %d) = %d, want %d", tt.stations, tt.r, tt.k, got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower2(tt.stations, tt.r, tt.k); got != tt.want {
				t.Fatalf("maxPower2(%v, %d, %d) = %d, want %d", tt.stations, tt.r, tt.k, got, tt.want)
			}
		})
	}
}
