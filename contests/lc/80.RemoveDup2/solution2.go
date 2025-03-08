package main

func removeDuplicates2(nums []int) int {
	j := 0
	i := 0
	l := len(nums)
	for j < l {
		if j < l-2 && nums[j] == nums[j+2] {
			j++
			continue
		}
		if j < l {
			nums[i] = nums[j]
			i++
			j++
		}
	}
	return i
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// func main() {
// 	nums1 := []int{1, 1, 1, 2, 2, 3}
// 	k1 := removeDuplicates(nums1)
// 	fmt.Println("Пример 1:")
// 	fmt.Println("Количество элементов:", k1)
// 	fmt.Println("Массив после модификации:", nums1[:k1])

// 	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
// 	k2 := removeDuplicates(nums2)
// 	fmt.Println("\nПример 2:")
// 	fmt.Println("Количество элементов:", k2)
// 	fmt.Println("Массив после модификации:", nums2[:k2])
// }
