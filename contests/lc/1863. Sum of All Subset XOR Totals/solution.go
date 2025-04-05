package main

import (
	"fmt"
)

// subsetXORSum вычисляет сумму XOR‑сумм всех подмножеств массива nums.
func subsetXORSum(nums []int) int {
	n := len(nums)
	orVal := 0
	// Вычисляем побитовое ИЛИ всех элементов.
	for _, num := range nums {
		orVal |= num
	}
	// Результат равен orVal * 2^(n-1). Умножение на 2^(n-1) можно выполнить сдвигом влево.
	return orVal << (n - 1)
}

func main() {
	// Пример 1:
	nums1 := []int{1, 3}
	fmt.Printf("nums = %v → сумма XOR подмножеств = %d\n", nums1, subsetXORSum(nums1)) // Ожидается 6

	// Пример 2:
	nums2 := []int{5, 1, 6}
	// Вычислим: orVal = 5 | 1 | 6 = 5 | 1 = 5, 5 | 6 = 7; n = 3, 2^(3-1)=4, ответ = 7*4 = 28.
	fmt.Printf("nums = %v → сумма XOR подмножеств = %d\n", nums2, subsetXORSum(nums2))
}
