# [2598. Smallest Missing Non-negative Integer After Operations](https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/description)

## Problem

You are given a 0-indexed integer array `nums` and an integer `value`. In one operation you may add or subtract `value` from any element of `nums`.

The MEX (minimum excluded) of an array is the smallest missing non-negative integer in it. Return the maximum MEX you can achieve after applying any number of the allowed operations.

## Examples

### Example 1

```
Input: nums = [1,-10,7,13,6,8], value = 5
Output: 4
```

Explanation: Apply the following operations:
- Add `value` to `nums[1]` twice → `nums = [1,0,7,13,6,8]`
- Subtract `value` from `nums[2]` once → `nums = [1,0,2,13,6,8]`
- Subtract `value` from `nums[3]` twice → `nums = [1,0,2,3,6,8]`

The resulting MEX is 4, which is the maximum achievable.

### Example 2

```
Input: nums = [1,-10,7,13,6,8], value = 7
Output: 2
```

Explanation: Subtract `value` from `nums[2]` once to obtain `nums = [1,-10,0,13,6,8]`. The MEX is 2, which is the maximum achievable.

## Constraints

- `1 <= nums.length, value <= 10^5`
- `-10^9 <= nums[i] <= 10^9`
