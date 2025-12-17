package main

import (
	"reflect"
	"testing"
)

func TestCountMentions(t *testing.T) {
	tests := []struct {
		name          string
		numberOfUsers int
		events        [][]string
		wantMentions  []int
	}{
		{
			name:          "example 1",
			numberOfUsers: 2,
			events: [][]string{
				{"MESSAGE", "10", "id1 id0"},
				{"OFFLINE", "11", "0"},
				{"MESSAGE", "71", "HERE"},
			},
			wantMentions: []int{2, 2},
		},
		{
			name:          "example 2",
			numberOfUsers: 2,
			events: [][]string{
				{"MESSAGE", "10", "id1 id0"},
				{"OFFLINE", "11", "0"},
				{"MESSAGE", "12", "ALL"},
			},
			wantMentions: []int{2, 2},
		},
		{
			name:          "example 3",
			numberOfUsers: 2,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "12", "HERE"},
			},
			wantMentions: []int{0, 1},
		},
		{
			name:          "offline then same timestamp message",
			numberOfUsers: 3,
			events: [][]string{
				{"OFFLINE", "5", "1"},
				{"MESSAGE", "5", "HERE"},
			},
			wantMentions: []int{1, 0, 1},
		},
		{
			name:          "duplicates and auto return",
			numberOfUsers: 2,
			events: [][]string{
				{"OFFLINE", "10", "0"},
				{"MESSAGE", "12", "id0 id0 id1"},
				{"MESSAGE", "70", "HERE"},
				{"MESSAGE", "71", "HERE"},
			},
			wantMentions: []int{4, 3},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := countMentions(tt.numberOfUsers, tt.events)
			if !reflect.DeepEqual(got, tt.wantMentions) {
				t.Fatalf("countMentions() = %v, want %v", got, tt.wantMentions)
			}
		})
	}
}
