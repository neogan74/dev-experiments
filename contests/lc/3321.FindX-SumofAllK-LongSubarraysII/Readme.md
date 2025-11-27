# [3321. Find X-Sum of All K-Long Subarrays II](https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-ii/description)

## Problem Statement
You are given an array `nums` of `n` integers and two integers `k` and `x`.

Return an integer array `answer` of length `n - k + 1` where `answer[i]` is the x-sum of the subarray `nums[i..i + k - 1]`.

## X-Sum Calculation
1. Count the occurrences of every element in the array.
2. Keep only the occurrences of the top `x` most frequent elements. When two elements have the same frequency, the larger value is considered more frequent.
3. Sum the remaining elements.

If the array has fewer than `x` distinct elements, the x-sum is simply the sum of the entire array.

## Examples
**Example 1**
```text
Input: nums = [1,1,2,2,3,4,2,3], k = 6, x = 2
Output: [6,10,12]
```
- For subarray `[1,1,2,2,3,4]`, keep elements `1` and `2`, so `answer[0] = 1 + 1 + 2 + 2`.
- For subarray `[1,2,2,3,4,2]`, keep elements `2` and `4`, so `answer[1] = 2 + 2 + 2 + 4`.
- For subarray `[2,2,3,4,2,3]`, keep elements `2` and `3`, so `answer[2] = 2 + 2 + 2 + 3 + 3`.

**Example 2**
```text
Input: nums = [3,8,7,8,7,5], k = 2, x = 2
Output: [11,15,15,15,12]
```
- Since `k == x`, each `answer[i]` equals the sum of the subarray `nums[i..i + k - 1]`.

## Constraints
- `nums.length == n`
- `1 <= n <= 10^5`
- `1 <= nums[i] <= 10^9`
- `1 <= x <= k <= nums.length`
