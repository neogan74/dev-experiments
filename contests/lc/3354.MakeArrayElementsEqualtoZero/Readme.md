# 3354. [Make Array Elements Equal to Zero](https://leetcode.com/problems/make-array-elements-equal-to-zero/description)

## Problem
You are given an integer array `nums`. Pick a starting index `curr` such that `nums[curr] == 0`, and choose an initial movement direction (`left` or `right`). Then repeat the following steps until you leave the array (`curr < 0` or `curr >= n`):

1. If `nums[curr] == 0`, move one step in the current direction.
2. If `nums[curr] > 0`:
   - Decrement `nums[curr]` by 1.
   - Reverse your direction.
   - Move one step in the new direction.

The selection `(curr, direction)` is **valid** if every element of `nums` becomes zero by the time the process ends. Return the number of valid selections.

## Examples
**Example 1**

- Input: `nums = [1,0,2,0,3]`
- Output: `2`
- Valid starting choices:

```
curr = 3, direction = left
[1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,1,0,3] -> [1,0,1,0,3] -> [1,0,1,0,2]
-> [1,0,1,0,2] -> [1,0,0,0,2] -> [1,0,0,0,2] -> [1,0,0,0,1] -> [1,0,0,0,1]
-> [1,0,0,0,1] -> [1,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1] -> [0,0,0,0,1]
-> [0,0,0,0,1] -> [0,0,0,0,0]

curr = 3, direction = right
[1,0,2,0,3] -> [1,0,2,0,3] -> [1,0,2,0,2] -> [1,0,2,0,2] -> [1,0,1,0,2]
-> [1,0,1,0,2] -> [1,0,1,0,1] -> [1,0,1,0,1] -> [1,0,0,0,1] -> [1,0,0,0,1]
-> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [1,0,0,0,0] -> [0,0,0,0,0]
```

**Example 2**

- Input: `nums = [2,3,4,0,4,1,0]`
- Output: `0`
- Explanation: No starting choice leads to all zeros.

## Constraints
- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`
- At least one index `i` satisfies `nums[i] == 0`
