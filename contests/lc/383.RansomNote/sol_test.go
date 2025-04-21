package _83_RansomNote

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	got := one()
	want := "coll"

	if got != want {
		_ = fmt.Errorf("got %s, want %s", got, want)
	}
}

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
		{"", "abc", true},
		{"abc", "", false},
		{"abc", "aabbcc", true},
	}

	for _, tt := range tests {
		result := canConstruct(tt.ransomNote, tt.magazine)
		if result != tt.expected {
			t.Errorf("canConstruct(%q, %q) = %v; want %v",
				tt.ransomNote, tt.magazine, result, tt.expected)
		}
	}
}
