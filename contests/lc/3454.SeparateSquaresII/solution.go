package main

import "sort"

type Event struct {
	y     int64
	x1    int64
	x2    int64
	delta int
}

type SegTree struct {
	xs     []int64
	cnt    []int
	length []int64
	n      int
}

func NewSegTree(xs []int64) *SegTree {
	n := len(xs) - 1
	size := 4 * n
	return &SegTree{
		xs:     xs,
		cnt:    make([]int, size),
		length: make([]int64, size),
		n:      n,
	}
}

func (st *SegTree) Update(node, l, r, ql, qr, delta int) {
	if ql > r || qr < l {
		return
	}
	if ql <= l && r <= qr {
		st.cnt[node] += delta
		st.pushUp(node, l, r)
		return
	}
	mid := (l + r) / 2
	st.Update(node*2, l, mid, ql, qr, delta)
	st.Update(node*2+1, mid+1, r, ql, qr, delta)
	st.pushUp(node, l, r)
}

func (st *SegTree) pushUp(node, l, r int) {
	if st.cnt[node] > 0 {
		st.length[node] = st.xs[r+1] - st.xs[l]
		return
	}
	if l == r {
		st.length[node] = 0
		return
	}
	st.length[node] = st.length[node*2] + st.length[node*2+1]
}

func separateSquares(squares [][]int) float64 {
	n := len(squares)
	events := make([]Event, 0, 2*n)
	xs := make([]int64, 0, 2*n)

	for _, s := range squares {
		x := int64(s[0])
		y := int64(s[1])
		l := int64(s[2])
		events = append(events, Event{y: y, x1: x, x2: x + l, delta: 1})
		events = append(events, Event{y: y + l, x1: x, x2: x + l, delta: -1})
		xs = append(xs, x, x+l)
	}

	sort.Slice(xs, func(i, j int) bool { return xs[i] < xs[j] })
	xs = unique(xs)

	sort.Slice(events, func(i, j int) bool {
		if events[i].y == events[j].y {
			return events[i].delta > events[j].delta
		}
		return events[i].y < events[j].y
	})

	st := NewSegTree(xs)
	total := sweepTotal(events, xs, st)
	target := total / 2.0

	st = NewSegTree(xs)
	return sweepFind(events, xs, st, target)
}

func sweepTotal(events []Event, xs []int64, st *SegTree) float64 {
	if len(events) == 0 {
		return 0
	}
	prevY := events[0].y
	area := 0.0

	for i := 0; i < len(events); {
		y := events[i].y
		dy := y - prevY
		if dy > 0 {
			area += float64(st.length[1]) * float64(dy)
		}
		for i < len(events) && events[i].y == y {
			l := lowerBound(xs, events[i].x1)
			r := lowerBound(xs, events[i].x2) - 1
			if l <= r {
				st.Update(1, 0, st.n-1, l, r, events[i].delta)
			}
			i++
		}
		prevY = y
	}
	return area
}

func sweepFind(events []Event, xs []int64, st *SegTree, target float64) float64 {
	if len(events) == 0 {
		return 0
	}
	prevY := events[0].y
	area := 0.0

	for i := 0; i < len(events); {
		y := events[i].y
		dy := y - prevY
		if dy > 0 {
			span := float64(st.length[1]) * float64(dy)
			if st.length[1] > 0 && area+span >= target {
				return float64(prevY) + (target-area)/float64(st.length[1])
			}
			area += span
		}
		for i < len(events) && events[i].y == y {
			l := lowerBound(xs, events[i].x1)
			r := lowerBound(xs, events[i].x2) - 1
			if l <= r {
				st.Update(1, 0, st.n-1, l, r, events[i].delta)
			}
			i++
		}
		prevY = y
	}
	return float64(prevY)
}

func lowerBound(xs []int64, x int64) int {
	return sort.Search(len(xs), func(i int) bool { return xs[i] >= x })
}

func unique(xs []int64) []int64 {
	if len(xs) == 0 {
		return xs
	}
	out := []int64{xs[0]}
	for i := 1; i < len(xs); i++ {
		if xs[i] != out[len(out)-1] {
			out = append(out, xs[i])
		}
	}
	return out
}
