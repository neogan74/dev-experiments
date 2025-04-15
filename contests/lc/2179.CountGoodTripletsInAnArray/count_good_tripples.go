package main

import "fmt"

type FenwickTree struct {
	tree []int
	size int
}

func NewFenwickTree(size int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, size+2),
		size: size + 2,
	}
}

func (ft *FenwickTree) Update(index, delta int) {
	index += 1
	for index < ft.size {
		ft.tree[index] += delta
		index += index & -index
	}
}

func (ft *FenwickTree) Query(index int) int {
	index += 1
	result := 0
	for index > 0 {
		result += ft.tree[index]
		index -= index & -index
	}
	return result
}

func countGoodTriplets(nums1, nums2 []int) int64 {
	n := len(nums1)
	pos := make([]int, n)
	for i, val := range nums1 {
		pos[val] = i
	}

	indexInNums1 := make([]int, n)
	for i, val := range nums2 {
		indexInNums1[i] = pos[val]
	}

	leftTree := NewFenwickTree(n)
	rightTree := NewFenwickTree(n)

	for _, idx := range indexInNums1 {
		rightTree.Update(idx, 1)
	}

	var result int64 = 0
	for _, idx := range indexInNums1 {
		rightTree.Update(idx, -1)
		leftCount := leftTree.Query(idx - 1)
		rightCount := rightTree.Query(n-1) - rightTree.Query(idx)
		result += int64(leftCount) * int64(rightCount)
		leftTree.Update(idx, 1)
	}

	return result
}

func main() {
	nums1 := []int{2, 0, 1, 3}
	nums2 := []int{0, 1, 2, 3}
	fmt.Println(countGoodTriplets(nums1, nums2)) // Output: 1
}
