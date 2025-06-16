package _3355_zero_array_tranform

func isZeroArray(nums []int, queries [][]int) bool {
	n := len(nums)
	diff := make([]int, n+1)

	for _, q := range queries {
		l, r := q[0], q[1]
		diff[l]++
		if r+1 < len(diff) {
			diff[r+1]--
		}
	}

	curr := 0
	for i := 0; i < n; i++ {
		curr += diff[i]
		if nums[i] > curr {
			return false
		}
	}
	return true
}
