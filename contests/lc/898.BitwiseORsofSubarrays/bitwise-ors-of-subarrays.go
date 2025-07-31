package bitwiseorsofsubarrays

func subarrayBitwiseORs(arr []int) int {
	ans := make(map[int]struct{})
	for i := 0; i < len(arr); i++ {
		ans[arr[i]] = struct{}{}
		for j := i - 1; j >= 0; j-- {
			if arr[j]|arr[i] == arr[j] {
				break
			}
			arr[j] |= arr[i]
			ans[arr[j]] = struct{}{}
		}
	}
	return len(ans)
}
