package _3362_zero_array_transformation_iii

import (
	"container/heap"
	"sort"
)

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	used := &minHeap{}
	available := &maxHeap{}
	heap.Init(used)
	heap.Init(available)

	queryPos := 0
	appliedCount := 0

	for i := 0; i < n; i++ {
		// Добавляем все запросы, начинающиеся в i
		for queryPos < len(queries) && queries[queryPos][0] == i {
			heap.Push(available, queries[queryPos][1])
			queryPos++
		}

		// Уменьшаем nums[i] на количество активных запросов
		nums[i] -= used.Len()

		// Применяем доступные запросы, если nums[i] > 0
		for nums[i] > 0 && available.Len() > 0 && (*available)[0] >= i {
			end := heap.Pop(available).(int)
			heap.Push(used, end)
			nums[i]--
			appliedCount++
		}

		// Если не удалось уменьшить nums[i] до 0
		if nums[i] > 0 {
			return -1
		}

		// Удаляем завершившиеся запросы
		for used.Len() > 0 && (*used)[0] == i {
			heap.Pop(used)
		}
	}

	return len(queries) - appliedCount
}
