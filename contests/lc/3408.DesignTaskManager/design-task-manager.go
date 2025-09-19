package designtaskmanager

import "container/heap"

type item struct {
	priority int
	taskId   int
}
type maxHeap []item

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool {
	// max-heap: выше приоритет -> выше, при равенстве — больше taskId -> выше
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].taskId > h[j].taskId
}
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)   { *h = append(*h, x.(item)) }
func (h *maxHeap) Pop() any     { old := *h; v := old[len(old)-1]; *h = old[:len(old)-1]; return v }
func (h *maxHeap) top() (item, bool) {
	if len(*h) == 0 {
		return item{}, false
	}
	return (*h)[0], true
}

type TaskManager struct {
	task2user map[int]int
	task2prio map[int]int
	h         maxHeap
}

func Constructor(tasks [][]int) TaskManager {
	tm := TaskManager{
		task2user: make(map[int]int),
		task2prio: make(map[int]int),
		h:         maxHeap{},
	}
	heap.Init(&tm.h)
	for _, t := range tasks {
		tm.Add(t[0], t[1], t[2])
	}
	return tm
}

func (tm *TaskManager) Add(userId, taskId, priority int) {
	tm.task2user[taskId] = userId
	tm.task2prio[taskId] = priority
	heap.Push(&tm.h, item{priority: priority, taskId: taskId})
}

func (tm *TaskManager) Edit(taskId, newPriority int) {
	// ленивая стратегия: просто обновим и положим новую запись
	tm.task2prio[taskId] = newPriority
	heap.Push(&tm.h, item{priority: newPriority, taskId: taskId})
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.task2user, taskId)
	delete(tm.task2prio, taskId)
}

func (tm *TaskManager) ExecTop() int {
	for tm.h.Len() > 0 {
		top := heap.Pop(&tm.h).(item)
		// валидна ли запись?
		if p, ok := tm.task2prio[top.taskId]; ok && p == top.priority {
			user := tm.task2user[top.taskId]
			delete(tm.task2user, top.taskId)
			delete(tm.task2prio, top.taskId)
			return user
		}
		// иначе протухшая — пропускаем
	}
	return -1
}
