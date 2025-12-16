package findingpairswithacertainsum

type FindSumPairs struct {
	nums1 []int
	nums2 []int
	cnt   map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	cnt := map[int]int{}
	for _, x := range nums2 {
		cnt[x]++
	}
	return FindSumPairs{nums1, nums2, cnt}
}

func (fsp *FindSumPairs) Add(index int, val int) {
	fsp.cnt[fsp.nums2[index]]--
	fsp.nums2[index] += val
	fsp.cnt[fsp.nums2[index]]++
}

func (fsp *FindSumPairs) Count(tot int) (ans int) {
	for _, x := range fsp.nums1 {
		ans += fsp.cnt[tot-x]
	}
	return
}

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
