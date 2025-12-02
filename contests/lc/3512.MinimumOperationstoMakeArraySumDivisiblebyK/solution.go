package _512_MinimumOperationstoMakeArraySumDivisiblebyK

// minOperations returns the smallest number of decrements needed
// to make the sum of nums divisible by k.
func minOperations(nums []int, k int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total % k
}
