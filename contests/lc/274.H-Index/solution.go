package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	for i := range citations {
		h := n - i // Number of papers with at least `citations[i]` citations
		if citations[i] >= h {
			return h // Maximum valid `h` found
		}
	}
	return 0
}

func main() {
	citations := []int{3, 0, 6, 1, 5}
	res := hIndex(citations)
	fmt.Printf("Result %v\n", res)
	citations2 := []int{1, 3, 1}
	res = hIndex(citations2)
	fmt.Printf("result %v\n", res)

}
