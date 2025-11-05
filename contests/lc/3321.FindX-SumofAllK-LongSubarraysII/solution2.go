package findxsumofallklongsubarraysii

import "container/heap"

type dhState uint8

const (
	dhNone dhState = iota
	dhHot
	dhRest
)

type dhRecord struct {
	val  int
	freq int
	idx  int
	st   dhState
}

type dhHotHeap []*dhRecord

func (h dhHotHeap) Len() int { return len(h) }
func (h dhHotHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq < h[j].freq
	}
	return h[i].val < h[j].val
}
func (h dhHotHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h *dhHotHeap) Push(x interface{}) {
	rec := x.(*dhRecord)
	rec.idx = len(*h)
	*h = append(*h, rec)
}
func (h *dhHotHeap) Pop() interface{} {
	old := *h
	rec := old[len(old)-1]
	*h = old[:len(old)-1]
	rec.idx = -1
	return rec
}

type dhRestHeap []*dhRecord

func (h dhRestHeap) Len() int { return len(h) }
func (h dhRestHeap) Less(i, j int) bool {
	if h[i].freq != h[j].freq {
		return h[i].freq > h[j].freq
	}
	return h[i].val > h[j].val
}
func (h dhRestHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}
func (h *dhRestHeap) Push(x interface{}) {
	rec := x.(*dhRecord)
	rec.idx = len(*h)
	*h = append(*h, rec)
}
func (h *dhRestHeap) Pop() interface{} {
	old := *h
	rec := old[len(old)-1]
	*h = old[:len(old)-1]
	rec.idx = -1
	return rec
}

func dhBetter(a, b *dhRecord) bool {
	if a.freq != b.freq {
		return a.freq > b.freq
	}
	return a.val > b.val
}

// findXSumTwoHeaps provides an alternative implementation using a dual-heap data structure.
// It keeps track of the x most frequent (larger value on ties) elements for the current window.
func findXSumTwoHeaps(nums []int, k int, x int) []int64 {
	n := len(nums)
	if n == 0 || k == 0 || x == 0 {
		return nil
	}

	result := make([]int64, n-k+1)
	hot := &dhHotHeap{}
	rest := &dhRestHeap{}
	heap.Init(hot)
	heap.Init(rest)

	freq := make(map[int]*dhRecord, k)
	var hotSum int64

	detach := func(rec *dhRecord) {
		switch rec.st {
		case dhHot:
			hotSum -= int64(rec.val) * int64(rec.freq)
			heap.Remove(hot, rec.idx)
			rec.st = dhNone
		case dhRest:
			heap.Remove(rest, rec.idx)
			rec.st = dhNone
		}
	}

	promote := func(rec *dhRecord) {
		rec.st = dhHot
		hotSum += int64(rec.val) * int64(rec.freq)
		heap.Push(hot, rec)
	}

	demote := func(rec *dhRecord) {
		rec.st = dhRest
		heap.Push(rest, rec)
	}

	rebuild := func() {
		for hot.Len() < x && rest.Len() > 0 {
			best := heap.Pop(rest).(*dhRecord)
			best.st = dhNone
			promote(best)
		}
		for hot.Len() > 0 && rest.Len() > 0 {
			worst := (*hot)[0]
			best := (*rest)[0]
			if !dhBetter(best, worst) {
				break
			}

			heap.Pop(hot)
			hotSum -= int64(worst.val) * int64(worst.freq)
			worst.st = dhNone

			heap.Pop(rest)
			best.st = dhNone

			demote(worst)
			promote(best)
		}
		for hot.Len() < x && rest.Len() > 0 {
			best := heap.Pop(rest).(*dhRecord)
			best.st = dhNone
			promote(best)
		}
	}

	add := func(v int) {
		rec, ok := freq[v]
		if !ok {
			rec = &dhRecord{val: v, idx: -1, st: dhNone}
			freq[v] = rec
		} else {
			detach(rec)
		}
		rec.freq++
		promote(rec)
		if hot.Len() > x {
			worst := heap.Pop(hot).(*dhRecord)
			hotSum -= int64(worst.val) * int64(worst.freq)
			worst.st = dhNone
			demote(worst)
		}
		rebuild()
	}

	erase := func(v int) {
		rec := freq[v]
		detach(rec)
		rec.freq--
		if rec.freq == 0 {
			delete(freq, v)
			rec.st = dhNone
			rec.idx = -1
			return
		}
		demote(rec)
		rebuild()
	}

	for i := 0; i < k; i++ {
		add(nums[i])
	}
	result[0] = hotSum

	for i := k; i < n; i++ {
		erase(nums[i-k])
		add(nums[i])
		result[i-k+1] = hotSum
	}
	return result
}
