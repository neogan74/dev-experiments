func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make([][2]int, len(profits))
	for i := range profits {
		projects[i] = [2]int{capital[i], profits[i]}
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i][0] < projects[j][0]
	})

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	ptr := 0

	for i := 0; i < k; i++ {
		for ptr < len(projects) && projects[ptr][0] <= w {
			heap.Push(maxHeap, projects[ptr][1])
			ptr++
		}
		if maxHeap.Len() == 0 {
			break
		}
		w += heap.Pop(maxHeap).(int)
	}

	return w
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
