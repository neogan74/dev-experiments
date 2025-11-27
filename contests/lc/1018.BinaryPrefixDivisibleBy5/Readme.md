# [1018. Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/description)

## Problem
Given a binary array `nums`, define `x_i` as the number whose binary representation is the prefix `nums[0..i]` (most-significant bit first). Return a boolean array `answer` where `answer[i]` is `true` if `x_i` is divisible by 5.

## Examples
- Example 1  
  Input: `nums = [0,1,1]`  
  Output: `[true,false,false]`  
  Explanation: Prefix values are 0 (divisible by 5), 1, 3.

- Example 2  
  Input: `nums = [1,1,1]`  
  Output: `[false,false,false]`

## Constraints
- `1 <= nums.length <= 10^5`
- `nums[i]` is either `0` or `1`
