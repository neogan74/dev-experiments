package main

import "fmt"

func minOperations(nums []int) int {
	count := 0
	n := len(nums)

	for i := 0; i <= n-3; i++ {
		if nums[i] == 0 {
			// Инвертируем текущий и два следующих элемента
			nums[i] ^= 1
			nums[i+1] ^= 1
			nums[i+2] ^= 1
			count++
		}
	}

	// Проверим последние два элемента
	if nums[n-2] == 0 || nums[n-1] == 0 {
		return -1
	}

	return count
}

func main() {
	fmt.Println(minOperations([]int{0, 1, 1, 1, 0, 0})) // Ожидаем: 3
	fmt.Println(minOperations([]int{0, 1, 1, 1}))       // Ожидаем: -1
}
