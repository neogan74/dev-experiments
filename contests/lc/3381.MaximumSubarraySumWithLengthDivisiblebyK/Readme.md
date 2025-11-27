# [3381. Maximum Subarray Sum With Length Divisible by K](https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/description)

## Problem
Given an integer array `nums` and an integer `k`, return the maximum possible sum of a subarray whose length is divisible by `k`.

## Examples
**Example 1**
```
Input:  nums = [1, 2], k = 1
Output: 3
Explanation: Subarray [1, 2] has length 2, which is divisible by 1.
```

**Example 2**
```
Input:  nums = [-1, -2, -3, -4, -5], k = 4
Output: -10
Explanation: The best subarray is [-1, -2, -3, -4], length 4 divisible by 4.
```

**Example 3**
```
Input:  nums = [-5, 1, 2, -3, 4], k = 2
Output: 4
Explanation: The best subarray is [1, 2, -3, 4], length 4 divisible by 2.
```

## Constraints
- `1 <= k <= nums.length <= 2 * 10^5`
- `-10^9 <= nums[i] <= 10^9`
