package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeGrid(t *testing.T, dir string, lines []string) string {
	t.Helper()
	p := filepath.Join(dir, "grid.txt")
	err := os.WriteFile(p, []byte(joinLines(lines)), 0644)
	if err != nil {
		t.Fatalf("write grid: %v", err)
	}
	return p
}

func joinLines(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	out := lines[0]
	for _, l := range lines[1:] {
		out += "\n" + l
	}
	return out
}

func loadGrid(t *testing.T, path string) []string {
	t.Helper()
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read grid: %v", err)
	}
	return splitLines(string(raw))
}

func splitLines(s string) []string {
	if s == "" {
		return nil
	}
	out := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	if start <= len(s) {
		out = append(out, s[start:])
	}
	return out
}

func TestEmpty(t *testing.T) {
	tmp := t.TempDir()
	grid := writeGrid(t, tmp, nil)
	lines := loadGrid(t, grid)
	if got := countSplits(lines); got != 0 {
		t.Fatalf("want 0 splits, got %d", got)
	}
	if got := countTimelines(lines).String(); got != "0" {
		t.Fatalf("want 0 timelines, got %s", got)
	}
}

func TestNoStart(t *testing.T) {
	lines := []string{"....", ".^.."}
	if got := countSplits(lines); got != 0 {
		t.Fatalf("want 0 splits, got %d", got)
	}
	if got := countTimelines(lines).String(); got != "0" {
		t.Fatalf("want 0 timelines, got %s", got)
	}
}

func TestSingleSplit(t *testing.T) {
	lines := []string{"..S..", "..^.."}
	if got := countSplits(lines); got != 1 {
		t.Fatalf("want 1 split, got %d", got)
	}
	if got := countTimelines(lines).String(); got != "2" {
		t.Fatalf("want 2 timelines, got %s", got)
	}
}

func TestEdgeFalloff(t *testing.T) {
	lines := []string{"S", "^"}
	if got := countSplits(lines); got != 1 {
		t.Fatalf("want 1 split, got %d", got)
	}
	if got := countTimelines(lines).String(); got != "0" {
		t.Fatalf("want 0 timelines, got %s", got)
	}
}

func TestSample(t *testing.T) {
	lines := loadGrid(t, "input.txt")
	if got := countSplits(lines); got != 21 {
		t.Fatalf("want 21 splits, got %d", got)
	}
	if got := countTimelines(lines).String(); got != "40" {
		t.Fatalf("want 40 timelines, got %s", got)
	}
}
