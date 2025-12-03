package closestprimenumbersinrange

import "math"

func closestPrimes2(left int, right int) []int {
	isNotPrime := make([]bool, right+1)
	isNotPrime[0] = true
	isNotPrime[1] = true
	upperBound := int(math.Sqrt(float64(right)))
	for num := 2; num <= upperBound; num++ {
		if !isNotPrime[num] {
			for j := num * num; j <= right; j += num {
				isNotPrime[j] = true
			}
		}
	}

	prev := -1
	curMin := right
	ans := []int{-1, -1}
	for i := left; i < right+1; i++ {
		if !isNotPrime[i] {
			if prev != -1 && curMin > i-prev {
				curMin = i - prev
				ans = []int{prev, i}
			}
			prev = i
		}
	}
	return ans
}
