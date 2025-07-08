package findingpairswithacertainsum

import "testing"

func TestFindSumPais(t *testing.T) {
	obj := Constructor([]int{1, 2, 3}, []int{3, 4})
	// Test 1
	got := obj.Count(5)
	want := 2

	if got != want {
		t.Errorf("Count(5) = %d, want %d", got, want)
	}
	// Perform Add(0,1): nums2 becomes [4,4]
	obj.Add(0, 1)

	//Test 2
	got = obj.Count(6)
	want = 2
	if got != want {
		t.Errorf("Count(6) = %d, want %d", got, want)
	}
	// Test 3
	got = obj.Count(10)
	want = 0
	if got != want {
		t.Errorf("Count(6) = %d, want %d", got, want)
	}
}
