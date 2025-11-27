package main

// prefixesDivBy5 returns a boolean slice where each element indicates whether
// the corresponding binary prefix of nums is divisible by 5. We keep only the
// value modulo 5 while iterating to avoid overflow and unnecessary large ints.
func prefixesDivBy5(nums []int) []bool {
	result := make([]bool, len(nums))
	mod := 0

	for i, bit := range nums {
		mod = ((mod << 1) + bit) % 5
		result[i] = mod == 0
	}

	return result
}
