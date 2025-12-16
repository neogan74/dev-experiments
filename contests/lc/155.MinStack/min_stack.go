package _55_MinStack

import "math"

type MinStack struct {
	stk1 []int
	stk2 []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{math.MaxInt32}}
}

func (st *MinStack) Push(val int) {
	st.stk1 = append(st.stk1, val)
	st.stk2 = append(st.stk2, min(val, st.stk2[len(st.stk2)-1]))
}

func (st *MinStack) Pop() {
	st.stk1 = st.stk1[:len(st.stk1)-1]
	st.stk2 = st.stk2[:len(st.stk2)-1]
}

func (st *MinStack) Top() int {
	return st.stk1[len(st.stk1)-1]
}

func (st *MinStack) GetMin() int {
	return st.stk2[len(st.stk2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
