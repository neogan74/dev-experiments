# [3578. Count Partitions With Max-Min Difference at Most K](https://leetcode.com/problems/count-partitions-with-max-min-difference-at-most-k/description)

Partition the array into non-empty contiguous segments so that, in every segment, `max - min <= k`. Return the number of valid partitions modulo `1e9 + 7`.

## Approach
- Maintain a sliding window `[l, r]` where the current segment max/min are tracked with two monotonic queues. Shrink `l` while `max - min > k`; then every segment ending at `r` must start at `>= l`.
- Let `dp[i]` be the number of ways to partition the prefix `nums[:i]` (1-indexed positions). For position `i`, the last segment can start at any `j` with `l-1 <= j <= i-1`, so  
  `dp[i] = sum(dp[j]) for j in [l-1, i-1]`.
- Keep a running prefix sum of `dp` values to compute this range sum in O(1). Initialize `dp[0] = 1` for the empty prefix.

## Correctness Sketch
For each `i`, the window invariants guarantee any start index before `l` would violate `max - min <= k`, while any start at or after `l` is valid. Summing `dp[j]` over these valid starts counts all partitions of the prefix ending at `i`. Prefix sums return this total in constant time, so `dp[i]` exactly captures the valid partitions, and the recurrence builds the answer for the entire array.

## Complexity
- Time: O(n) using monotonic queues and prefix sums.
- Space: O(n) for `dp`/prefix arrays; the queues are O(n) in the worst case but O(1) amortized per step.

## Examples
- `nums = [9,4,1,3,7], k = 4` → `6`
- `nums = [3,3,4], k = 0` → `2`
