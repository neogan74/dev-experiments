package main

import (
	"fmt"
)

const MOD = int64(1e9 + 7)

func power(x, n, mod int64) int64 {
	res := int64(1)
	x = x % mod
	for n > 0 {
		if n%2 == 1 {
			res = (res * x) % mod
		}
		x = (x * x) % mod
		n /= 2
	}
	return res
}

func countGoodNumbers(n int64) int {
	even := (n + 1) / 2
	odd := n / 2

	result := (power(5, odd, MOD) * power(4, even, MOD)) % MOD
	return int(result)
}

func main() {
	var n int64
	fmt.Print("Enter n: ")
	fmt.Scan(&n)

	fmt.Println(countGoodNumbers(n))
}
