package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// Сортируем массив.
	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// Пропускаем дубликаты для первого элемента.
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				// Добавляем найденную тройку.
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// Пропускаем дубликаты для left.
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Пропускаем дубликаты для right.
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else { // sum > 0
				right--
			}
		}
	}

	return res
}

func main() {
	// Пример:
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println("Найденные тройки:")
	for _, triplet := range result {
		fmt.Println(triplet)
	}
	// Ожидаемые тройки: [-1, -1, 2] и [-1, 0, 1]
}
