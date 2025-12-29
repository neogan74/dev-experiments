package main

import "testing"

func TestPyramidTransition(t *testing.T) {
	cases := []struct {
		name    string
		bottom  string
		allowed []string
		want    bool
	}{
		{
			name:    "example-true",
			bottom:  "BCD",
			allowed: []string{"BCC", "CDE", "CEA", "FFF"},
			want:    true,
		},
		{
			name:    "example-false",
			bottom:  "AAAA",
			allowed: []string{"AAB", "AAC", "BCD", "BBE", "DEF"},
			want:    false,
		},
		{
			name:    "single-step",
			bottom:  "AB",
			allowed: []string{"ABA"},
			want:    true,
		},
		{
			name:    "no-rule",
			bottom:  "AB",
			allowed: []string{"BCD"},
			want:    false,
		},
		{
			name:    "branching-success",
			bottom:  "ABC",
			allowed: []string{"ABD", "BCE", "BDF", "CEF"},
			want:    true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := pyramidTransition(tc.bottom, tc.allowed); got != tc.want {
				t.Fatalf("pyramidTransition(%q, %v) = %v, want %v", tc.bottom, tc.allowed, got, tc.want)
			}
		})
	}
}
