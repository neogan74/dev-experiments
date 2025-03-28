package main

// Greedily increment k while going through each number

func minZeroArray2(nums []int, queries [][]int) int {
	increments := make([]int, len(nums)+1)
	k := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		for sum+increments[i] < nums[i] {
			if k >= len(queries) { // running out of queries
				return -1
			}
			query := queries[k]
			l := query[0]
			r := query[1]
			val := query[2]
			// if r < i, the query doesn't improve increments[i], so just keep going to the next query
			if r >= i {
				// if l < i, we mark increments[i] so that it's effective immediately
				increments[max(l, i)] += val
				// subtract it down when getting outside of the range
				increments[r+1] -= val
			}
			k++
		}
		sum += increments[i]
	}
	return k
}

/*
// Binary search + line sweep, O(NlogM), where N = len(nums) and M = len(queries)

func minZeroArray(nums []int, queries [][]int) int {
    increments := make([]int, len(nums))
    // check whether the nums is already a zero array
    if canZero(0, 0, -1, true, increments, nums, queries) {
        return 0
    }
    l, r := 0, len(queries)
    isRightHalf := true
    for l < r {
        m := l + (r - l) / 2
        if canZero(l, r, m, isRightHalf, increments, nums, queries) {
            r = m
            isRightHalf = false
        } else {
            l = m + 1
            isRightHalf = true
        }
    }
    if l + 1 > len(queries) {
        return -1
    } else {
        return l + 1
    }
}

func canZero(l, r, m int, isRightHalf bool, increments []int, nums []int, queries [][]int) bool {
    var s, e, multiplier int
    if isRightHalf {
        // add it from l to m (inclusive) to the increments
        s = l
        e = m + 1
        multiplier = 1
    } else {
        // subtract it from m + 1 to r (inclusive) from the increments
        s = m + 1
        e = r + 1
        multiplier = -1
    }
    for i := s; i < e; i++ {
        query := queries[i]
        increments[query[0]] += multiplier * query[2]
        if query[1] + 1 < len(nums) {
            increments[query[1] + 1] -= multiplier * query[2]
        }
    }
    // get a running sum of increments and compare it with nums
    runningSum := 0
    for i := 0; i < len(nums); i++ {
        runningSum += increments[i]
        if runningSum < nums[i] {
            return false
        }
    }
    return true
}
*/
