package _28_SummaryRanges

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	// Для корректного тестирования возвращаем непустой пустой срез (не nil) для пустого входа.
	if len(nums) == 0 {
		return []string{}
	}

	res := make([]string, 0)
	n := len(nums)
	i := 0
	for i < n {
		start := nums[i]
		j := i
		// Продвигаем j, пока следующий элемент = текущий + 1
		for j+1 < n && nums[j+1] == nums[j]+1 {
			j++
		}
		// Формируем строку диапазона
		if start == nums[j] {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[j]))
		}
		i = j + 1
	}
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))    // ["0->2","4->5","7"]
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9})) // ["0","2->4","6","8->9"]
}
