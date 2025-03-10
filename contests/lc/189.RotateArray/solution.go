package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n // Если k больше длины массива

	// Функция для разворота части массива
	reverse := func(nums []int, start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}

	// 1. Разворот всего массива
	reverse(nums, 0, n-1)
	// 2. Разворот первых k элементов
	reverse(nums, 0, k-1)
	// 3. Разворот оставшейся части массива
	reverse(nums, k, n-1)
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums1, 3)
	fmt.Println("После поворота:", nums1) // Вывод: [5,6,7,1,2,3,4]

	nums2 := []int{-1, -100, 3, 99}
	rotate(nums2, 2)
	fmt.Println("После поворота:", nums2) // Вывод: [3,99,-1,-100]
}
