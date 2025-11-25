package main

// smallestRepunitDivByK returns the length of the smallest repunit (all digits are 1)
// divisible by k, or -1 if none exists. For k divisible by 2 or 5, no such repunit
// can be divisible. Otherwise, at most k different remainders exist, so iterating
// lengths up to k guarantees finding the answer if it exists.
func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}

	remainder := 0
	for length := 1; length <= k; length++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return length
		}
	}

	return -1
}
