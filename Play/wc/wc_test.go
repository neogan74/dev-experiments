package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("Word1 word2 word3 word4\n")
	want := 4

	got := count(b, false, false)
	if got != want {
		t.Errorf("Expected %d, got %d instead.\n", want, got)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2\n line2 word3\n line3 word4\n")
	want := 3

	got := count(b, true, false)
	if want != got {
		t.Errorf("Expected %d, but got %d lines instead.", want, got)
	}
}
