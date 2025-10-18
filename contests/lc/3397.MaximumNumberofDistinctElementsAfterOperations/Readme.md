# [3397. Maximum Number of Distinct Elements After Operations](https://leetcode.com/problems/maximum-number-of-distinct-elements-after-operations/description)

## Problem
- You are given an integer array `nums` and an integer `k`.
- For each element, you may add an integer in the range `[-k, k]` at most once.
- Determine the maximum possible number of distinct elements after applying the operations.

## Examples
### Example 1
```text
Input:  nums = [1,2,2,3,3,4], k = 2
Output: 6
```
Applying the operation to the first four elements allows `nums` to become `[-1, 0, 1, 2, 3, 4]`, yielding six distinct values.

### Example 2
```text
Input:  nums = [4,4,4,4], k = 1
Output: 3
```
Adding `-1` to `nums[0]` and `1` to `nums[1]` transforms the array into `[3, 5, 4, 4]`, leaving three distinct values.

## Constraints
- `1 <= nums.length <= 10^5`
- `1 <= nums[i] <= 10^9`
- `0 <= k <= 10^9`
