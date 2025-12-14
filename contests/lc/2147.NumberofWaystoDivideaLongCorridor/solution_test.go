package numberofwaystodividealongcorridor

import "testing"

func TestNumberOfWays(t *testing.T) {
	tests := []struct {
		name     string
		corridor string
		want     int
	}{
		{name: "sample 1", corridor: "SSPPSPS", want: 3},
		{name: "sample 2", corridor: "PPSPSP", want: 1},
		{name: "sample 3", corridor: "S", want: 0},
		{name: "multiple options", corridor: "SSPPSS", want: 3},
		{name: "even seats but blocked", corridor: "SPPSPSSPPS", want: 0},
		{name: "no seats", corridor: "PPP", want: 0},
		{name: "single section", corridor: "SS", want: 1},
		{name: "odd seats overall", corridor: "PSSPPSPSPPSSPSP", want: 0},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfWays(tt.corridor); got != tt.want {
				t.Fatalf("numberOfWays(%q) = %d, want %d", tt.corridor, got, tt.want)
			}
		})
	}
}
