package vowelsgameinastring

import "testing"

func TestDoesAliceWin(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Test 1", "leetcoder", true},
		{"Test 2", "bbcd", false},
	}
	for _, tt := range tests {
		got := doesAliceWin(tt.s)
		if tt.expected != got {
			t.Errorf("doesAliceWin retunrs %v, but expected %v", got, tt.expected)
		}
	}
}
