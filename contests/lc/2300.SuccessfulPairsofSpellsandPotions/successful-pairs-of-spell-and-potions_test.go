package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSuccessfulPairs(t *testing.T) {
	cases := []struct {
		name    string
		spells  []int
		potions []int
		success int64
		want    []int
	}{
		{
			name:    "example 1",
			spells:  []int{5, 1, 3},
			potions: []int{1, 2, 3, 4, 5},
			success: 7,
			want:    []int{4, 0, 3},
		},
		{
			name:    "example 2",
			spells:  []int{3, 1, 2},
			potions: []int{8, 5, 8},
			success: 16,
			want:    []int{2, 0, 2},
		},
		{
			name:    "large success threshold",
			spells:  []int{1, 2, 3},
			potions: []int{1, 2, 3},
			success: 100,
			want:    []int{0, 0, 0},
		},
		{
			name:    "single element",
			spells:  []int{10},
			potions: []int{1},
			success: 5,
			want:    []int{1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := successfulPairs(tc.spells, tc.potions, tc.success)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("successfulPairs() = %v, want %v", got, tc.want)
			}
		})
	}
}

// bruteSuccessfulPairs is an O(n*m) reference implementation for validation.
func bruteSuccessfulPairs(spells []int, potions []int, success int64) []int {
	result := make([]int, len(spells))
	for i, spell := range spells {
		total := 0
		for _, potion := range potions {
			if int64(spell)*int64(potion) >= success {
				total++
			}
		}
		result[i] = total
	}
	return result
}

func TestSuccessfulPairsRandomized(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	for iter := 0; iter < 200; iter++ {
		n := r.Intn(8) + 1 // 1..8 spells
		m := r.Intn(8) + 1 // 1..8 potions
		spells := make([]int, n)
		potions := make([]int, m)
		for i := range spells {
			spells[i] = r.Intn(20) + 1 // 1..20
		}
		for i := range potions {
			potions[i] = r.Intn(20) + 1 // 1..20
		}
		success := int64(r.Intn(400) + 1)

		got := successfulPairs(append([]int(nil), spells...), append([]int(nil), potions...), success)
		want := bruteSuccessfulPairs(spells, potions, success)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("randomized case %d failed: spells=%v potions=%v success=%d got=%v want=%v", iter, spells, potions, success, got, want)
		}
	}
}

func BenchmarkSuccessfulPairs(b *testing.B) {
	spells := make([]int, 1000)
	potions := make([]int, 1000)
	for i := range spells {
		spells[i] = i%100 + 1
	}
	for i := range potions {
		potions[i] = (i*3)%100 + 1
	}
	success := int64(2500)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		successfulPairs(spells, potions, success)
	}
}
