package main

import (
	"fmt"
	"strconv"
)

func countSymmetricIntegers(low int, high int) int {
	count := 0
	for n := low; n <= high; n++ {
		s := strconv.Itoa(n)
		if len(s)%2 != 0 {
			continue
		}
		half := len(s) / 2
		left, right := 0, 0
		for i := 0; i < half; i++ {
			left += int(s[i] - '0')
			right += int(s[i+half] - '0')
		}
		if left == right {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countSymmetricIntegers(1, 1000))  // Пример: 9, 22, 33, 44, ..., 99, 1010 и т.д.
	fmt.Println(countSymmetricIntegers(10, 99))   // 9 чисел: 11, 22, ..., 99
	fmt.Println(countSymmetricIntegers(100, 300)) // Например: 121, 132, 213, 222, ...
}
