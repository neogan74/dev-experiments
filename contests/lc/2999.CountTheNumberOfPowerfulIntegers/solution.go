package main

import (
	"fmt"
	"strings"
)

func countPowerfulIntegers(start, finish int64, limit int, s string) int64 {
	low := fmt.Sprintf("%d", start)
	high := fmt.Sprintf("%d", finish)
	n := len(high)
	low = strings.Repeat("0", n-len(low)) + low // align digits
	pre_len := n - len(s)                       // prefix length
	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int, bool, bool) int64
	dfs = func(i int, limit_low, limit_high bool) int64 {
		// recursive boundary
		if i == n {
			return 1
		}
		if !limit_low && !limit_high && memo[i] != -1 {
			return memo[i]
		}
		lo := 0
		if limit_low {
			lo = int(low[i] - '0')
		}
		hi := 9
		if limit_high {
			hi = int(high[i] - '0')
		}

		var res int64 = 0
		if i < pre_len {
			for digit := lo; digit <= min(hi, limit); digit++ {
				res += dfs(i+1, limit_low && digit == lo, limit_high && digit == hi)
			}
		} else {
			x := int(s[i-pre_len] - '0')
			if lo <= x && x <= min(hi, limit) {
				res = dfs(i+1, limit_low && x == lo, limit_high && x == hi)
			}
		}

		if !limit_low && !limit_high {
			memo[i] = res
		}
		return res
	}
	return dfs(0, true, true)
}

func main() {
	start := int64(1)
	finish := int64(6000)
	limit := 4
	s := "124"

	result := countPowerfulIntegers(start, finish, limit, s)
	fmt.Println("Количество мощных чисел:", result)

	fmt.Println("Новый кейс")
	start = int64(20)
	finish = int64(1159)
	limit = 5
	s = "20"

	result = countPowerfulIntegers(start, finish, limit, s)
	fmt.Printf("Количество мощных чисел2: %d\n", result)

	fmt.Println("Новый кейс долгий")
	start = int64(1829505)
	finish = int64(1255574165)
	limit = 7
	s = "11223"

	result = countPowerfulIntegers(start, finish, limit, s)
	fmt.Printf("Количество мощных чисел(долгий): %d\n", result)
}
