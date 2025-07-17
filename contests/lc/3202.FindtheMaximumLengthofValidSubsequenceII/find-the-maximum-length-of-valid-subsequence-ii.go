package findthemaximumlengthofvalidsubsequenceii

func maximumLength(nums []int, k int) int {
	ans := 0
	f := make([][]int, k)
	for i := range f {
		f[i] = make([]int, k)
	}
	for _, x := range nums {
		x %= k
		for j := 0; j < k; j++ {
			y := (j - x + k) % k
			f[x][y] = f[y][x] + 1
			ans = max(ans, f[x][y])
		}
	}
	return ans
}
