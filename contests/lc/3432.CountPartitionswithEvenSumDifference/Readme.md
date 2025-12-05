# [3432. Count Partitions with Even Sum Difference](https://leetcode.com/problems/count-partitions-with-even-sum-difference/description)

You are given an integer array `nums` of length `n`.

A **partition** is any index `i` where `0 <= i < n - 1`, splitting the array into two non-empty subarrays:
- Left subarray: indices `[0, i]`
- Right subarray: indices `[i + 1, n - 1]`

Return the number of partitions where the difference between the sums of the left and right subarrays is even.

## Examples
**Example 1**  
Input: `nums = [10,10,3,7,6]`  
Output: `4`  
The valid partitions are:
- `[10] | [10, 3, 7, 6]` with sum difference `10 - 26 = -16` (even)
- `[10, 10] | [3, 7, 6]` with sum difference `20 - 16 = 4` (even)
- `[10, 10, 3] | [7, 6]` with sum difference `23 - 13 = 10` (even)
- `[10, 10, 3, 7] | [6]` with sum difference `30 - 6 = 24` (even)

**Example 2**  
Input: `nums = [1,2,2]`  
Output: `0`  
No partition results in an even difference.

**Example 3**  
Input: `nums = [2,4,6,8]`  
Output: `3`  
All partitions result in an even difference.

## Constraints
- `2 <= n == nums.length <= 100`
- `1 <= nums[i] <= 100`
