package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	// Если массив пустой, возвращаем 0
	if len(nums) == 0 {
		return 0
	}

	// i - индекс для записи следующего допустимого элемента
	i := 0
	// count - счётчик повторов
	count := 0

	// Проходим по всем элементам массива
	for j := 0; j < len(nums); j++ {
		// Если это первый элемент или элемент отличается от предыдущего
		if j == 0 || nums[j] != nums[j-1] {
			count = 1 // начинаем новый подсчёт
		} else {
			count++ // увеличиваем счётчик
		}

		// Если количество повторов не превышает 2, записываем элемент на позицию i
		if count <= 2 {
			nums[i] = nums[j]
			i++
		}
	}

	return i
}

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := removeDuplicates(nums1)
	k22 := removeDuplicates2(nums1)
	fmt.Println("Пример 1:")
	fmt.Println("Количество элементов:", k1)
	fmt.Println("Массив после модификации:", nums1[:k1])
	fmt.Println("Массив после модификации2:", nums1[:k22])

	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k2 := removeDuplicates(nums2)
	k3 := removeDuplicates2(nums2)
	fmt.Println("\nПример 2:")
	fmt.Println("Количество элементов:", k2)
	fmt.Println("Массив после модификации:", nums2[:k2])
	fmt.Println("Массив после модификации2:", nums2[:k3])
}
