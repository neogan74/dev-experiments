package main

import "fmt"

func fib(N int) int {
	if N <= 1 {
		return N
	}
	current, prev1, prev2 := 0, 1, 0
	for i := 2; i <= N; i++ {
		current = prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	return current
}

func main() {
	n := 3
	res := fib(n)
	fmt.Println(res)
}
