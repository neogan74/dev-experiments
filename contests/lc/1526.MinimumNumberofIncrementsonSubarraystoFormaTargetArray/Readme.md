# [1526. Minimum Number of Increments on Subarrays to Form a Target Array](https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/description)

## Problem

You are given an integer array `target`. You also have an integer array `initial` of the same size as `target` with all elements set to zero. In one operation, you can choose any subarray of `initial` and increment each value within that subarray by one. Return the minimum number of operations needed to transform `initial` into `target`. The test cases are generated so the answer fits in a 32-bit integer.

## Examples

### Example 1

Input: `target = [1,2,3,2,1]`  
Output: `3`

Explanation:
```
[0,0,0,0,0]  -> increment 1 from indices 0 to 4
[1,1,1,1,1]  -> increment 1 from indices 1 to 3
[1,2,2,2,1]  -> increment 1 at index 2
[1,2,3,2,1]  -> target formed
```

### Example 2

Input: `target = [3,1,1,2]`  
Output: `4`

Explanation:
```
[0,0,0,0] -> [1,1,1,1] -> [1,1,1,2] -> [2,1,1,2] -> [3,1,1,2]
```

### Example 3

Input: `target = [3,1,5,4,2]`  
Output: `7`

Explanation:
```
[0,0,0,0,0] -> [1,1,1,1,1] -> [2,1,1,1,1] -> [3,1,1,1,1]
           -> [3,1,2,2,2] -> [3,1,3,3,2] -> [3,1,4,4,2] -> [3,1,5,4,2]
```

## Constraints

- `1 <= target.length <= 10^5`
- `1 <= target[i] <= 10^5`
