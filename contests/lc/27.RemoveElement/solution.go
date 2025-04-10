package main

import "fmt"

func removeElement(nums []int, val int) []int {
	k := 0
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return nums
}

func main() {
	fmt.Println("First")
	nums := []int{3, 2, 2, 3}
	val := 3
	res := removeElement(nums, val)
	fmt.Println("Результат: ", res)
}
