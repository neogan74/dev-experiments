package _962_CountSubarraysWhereMaxElementAppearsAtLeastKTimes

import "fmt"

func countSubarrays(nums []int, k int) int64 {
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	var count, result int64
	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == maxNum {
			count++
		}

		for count >= int64(k) {
			// подсчёт всех подходящих подмассивов справа
			result += int64(len(nums) - right)
			if nums[left] == maxNum {
				count--
			}
			left++
		}
	}

	return result
}

func main() {
	fmt.Println(countSubarrays([]int{1, 3, 2, 3, 3}, 2)) // 6
	fmt.Println(countSubarrays([]int{1, 4, 2, 1}, 3))    // 0
}
