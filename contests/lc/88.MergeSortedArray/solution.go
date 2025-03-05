package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	nums1Copy := append([]int(nil), nums1[:m]...)
	p1 := 0
	p2 := 0
	for p := 0; p < m+n; p++ {
		if p2 >= n || (p1 < m && nums1Copy[p1] < nums2[p2]) {
			nums1[p] = nums1Copy[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
	}
	return nums1Copy
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	fmt.Println(merge(nums1, m, nums2, n))
}
