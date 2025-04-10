package main

import "fmt"

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	sum := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if i != num/i {
				sum += num / i
			}
		}
	}
	return sum == num
}

func main() {
	fmt.Println(checkPerfectNumber(28)) // true
	fmt.Println(checkPerfectNumber(6))  // true
	fmt.Println(checkPerfectNumber(27)) // false
	fmt.Println(checkPerfectNumber(1))  // false
}
