# [3350. Adjacent Increasing Subarrays Detection II](https://leetcode.com/problems/adjacent-increasing-subarrays-detection-ii/description)

## Problem

You are given an integer array `nums` of length `n`. Find the maximum value of `k` such that there exist two adjacent subarrays of length `k`, and both subarrays are strictly increasing.

The subarrays must satisfy all of the following:
- `nums[a..a + k - 1]` is strictly increasing.
- `nums[b..b + k - 1]` is strictly increasing.
- They are adjacent, meaning `b = a + k` with `a < b`.

Return the maximum possible value of `k`.

A subarray is a contiguous non-empty sequence of elements within an array.

## Examples

### Example 1
```
Input:  nums = [2, 5, 7, 8, 9, 2, 3, 4, 3, 1]
Output: 3
```
The subarray starting at index `2` is `[7, 8, 9]`, which is strictly increasing. The subarray starting at index `5` is `[2, 3, 4]`, which is also strictly increasing. These subarrays are adjacent, so the maximum valid `k` is `3`.

### Example 2
```
Input:  nums = [1, 2, 3, 4, 4, 4, 4, 5, 6, 7]
Output: 2
```
The subarray starting at index `0` is `[1, 2]`, which is strictly increasing. The subarray starting at index `2` is `[3, 4]`, which is also strictly increasing. These subarrays are adjacent, giving a maximum `k` of `2`.

## Constraints

- `2 <= nums.length <= 2 * 10^5`
- `-10^9 <= nums[i] <= 10^9`
