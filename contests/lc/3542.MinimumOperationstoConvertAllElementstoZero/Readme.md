# [3542. Minimum Operations to Convert All Elements to Zero](https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero/description)

You are given an array `nums` of size `n`, consisting of non-negative integers. Apply some (possibly zero) operations so that all elements become `0`.

In one operation, you can select a subarray `[i, j]` (`0 <= i <= j < n`) and set all occurrences of the minimum non-negative integer in that subarray to `0`.

Return the minimum number of operations required to make all elements in the array `0`.

## Examples

**Example 1**

```
Input: nums = [0,2]
Output: 1
```

Explanation: Select subarray `[1,1]` (which is `[2]`). The minimum value is `2`, so setting it to `0` makes the array `[0,0]`.

**Example 2**

```
Input: nums = [3,1,2,1]
Output: 3
```

Explanation:

1. Select `[1,3]` (`[1,2,1]`) ⇒ minimum `1` ⇒ array becomes `[3,0,2,0]`.
2. Select `[2,2]` (`[2]`) ⇒ minimum `2` ⇒ array becomes `[3,0,0,0]`.
3. Select `[0,0]` (`[3]`) ⇒ minimum `3` ⇒ array becomes `[0,0,0,0]`.

**Example 3**

```
Input: nums = [1,2,1,2,1,2]
Output: 4
```

Explanation:

1. Select `[0,5]` ⇒ minimum `1` ⇒ `[0,2,0,2,0,2]`.
2. Select `[1,1]` ⇒ minimum `2` ⇒ `[0,0,0,2,0,2]`.
3. Select `[3,3]` ⇒ minimum `2` ⇒ `[0,0,0,0,0,2]`.
4. Select `[5,5]` ⇒ minimum `2` ⇒ `[0,0,0,0,0,0]`.

## Constraints

- `1 <= n == nums.length <= 10^5`
- `0 <= nums[i] <= 10^5`
