package designmovierentalsystem

import (
	"reflect"
	"testing"
)

func TestMovieRentalWorkflow(t *testing.T) {
	system := Constructor(3, [][]int{
		{0, 1, 5},
		{0, 2, 6},
		{0, 3, 7},
		{1, 1, 4},
		{1, 2, 7},
		{2, 1, 5},
	})

	if got := system.Search(1); !reflect.DeepEqual(got, []int{1, 0, 2}) {
		t.Fatalf("initial search mismatch: got %v", got)
	}

	system.Rent(1, 1)
	if got := system.Search(1); !reflect.DeepEqual(got, []int{0, 2}) {
		t.Fatalf("search after first rent mismatch: got %v", got)
	}

	system.Rent(0, 1)
	if got := system.Search(1); !reflect.DeepEqual(got, []int{2}) {
		t.Fatalf("search after second rent mismatch: got %v", got)
	}

	expectedReport := [][]int{{1, 1}, {0, 1}}
	if got := system.Report(); !reflect.DeepEqual(got, expectedReport) {
		t.Fatalf("report mismatch: got %v", got)
	}

	if got := system.Report(); !reflect.DeepEqual(got, expectedReport) {
		t.Fatalf("report should be stable after peek: got %v", got)
	}

	system.Drop(1, 1)
	if got := system.Search(1); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Fatalf("search after drop mismatch: got %v", got)
	}

	if got := system.Report(); !reflect.DeepEqual(got, [][]int{{0, 1}}) {
		t.Fatalf("report after drop mismatch: got %v", got)
	}
}

func TestEmptyResults(t *testing.T) {
	system := Constructor(1, [][]int{{0, 0, 3}})

	if got := system.Search(1); got != nil {
		t.Fatalf("search for unknown movie expected nil, got %v", got)
	}

	system.Rent(0, 0)

	if got := system.Search(0); got != nil {
		t.Fatalf("search for rented movie expected nil, got %v", got)
	}

	if got := system.Report(); !reflect.DeepEqual(got, [][]int{{0, 0}}) {
		t.Fatalf("unexpected report result: %v", got)
	}

	system.Drop(0, 0)

	if got := system.Report(); got != nil {
		t.Fatalf("report after dropping all rentals expected nil, got %v", got)
	}
}
