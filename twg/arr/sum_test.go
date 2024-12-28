package arr

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestAllSumm(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{5, 6, 7}
	got := SumAll(s1, s2)
	want := []int{6, 18}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given %v and %v", got, want, s1, s2)
	}
}
