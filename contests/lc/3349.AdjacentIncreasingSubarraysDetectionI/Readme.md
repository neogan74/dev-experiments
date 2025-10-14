# [3349. Adjacent Increasing Subarrays Detection I](https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/description)

## Problem

Given an integer array `nums` of length `n` and an integer `k`, determine whether there exist two adjacent subarrays of length `k` such that both are strictly increasing. In other words, check if there are indices `a` and `b` (`a < b`) where:
- subarrays `nums[a..a + k - 1]` and `nums[b..b + k - 1]` are both strictly increasing
- the two subarrays are adjacent, meaning `b = a + k`

Return `true` if such subarrays exist, otherwise return `false`.

## Examples

### Example 1
```
Input:  nums = [2,5,7,8,9,2,3,4,3,1], k = 3
Output: true
```
Explanation: subarrays `[7,8,9]` (starting at index 2) and `[2,3,4]` (starting at index 5) are strictly increasing and adjacent.

### Example 2
```
Input:  nums = [1,2,3,4,4,4,4,5,6,7], k = 5
Output: false
```

## Constraints

- `2 <= nums.length <= 100`
- `1 < 2 * k <= nums.length`
- `-1000 <= nums[i] <= 1000`
