# [3512. Minimum Operations to Make Array Sum Divisible by K](https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/description/)

## Problem
- You are given an integer array `nums` and an integer `k`.
- In one operation, select an index `i` and replace `nums[i]` with `nums[i] - 1`.
- Return the minimum number of operations required so that the sum of the array is divisible by `k`.

## Examples
**Example 1**
```text
Input:  nums = [3,9,7], k = 5
Output: 4
```
Perform 4 operations on `nums[1] = 9`, obtaining `[3,5,7]` whose sum `15` is divisible by `5`.

**Example 2**
```text
Input:  nums = [4,1,3], k = 4
Output: 0
```
The sum `8` is already divisible by `4`.

**Example 3**
```text
Input:  nums = [3,2], k = 6
Output: 5
```
Decrease `nums[0]` three times and `nums[1]` twice to reach `[0,0]`, whose sum is divisible by `6`.

## Constraints
- `1 <= nums.length <= 1000`
- `1 <= nums[i] <= 1000`
- `1 <= k <= 100`
