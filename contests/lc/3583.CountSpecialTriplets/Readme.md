# [3583. Count Special Triplets](https://leetcode.com/problems/count-special-triplets/description)

Given an integer array `nums`, count the number of triplets `(i, j, k)` that satisfy:

- `0 <= i < j < k < n`, where `n = nums.length`
- `nums[i] = nums[j] * 2`
- `nums[k] = nums[j] * 2`

Return the total number of such triplets modulo `1_000_000_007`.

## Examples

- Input `nums = [6, 3, 6]` -> Output `1`  
  Triplet `(0, 1, 2)` with `6 = 3 * 2`.

- Input `nums = [0, 1, 0, 0]` -> Output `1`  
  Triplet `(0, 2, 3)` with `0 = 0 * 2`.

- Input `nums = [8, 4, 2, 8, 4]` -> Output `2`  
  Triplets `(0, 1, 3)` and `(1, 2, 4)` each satisfy the condition.

## Constraints

- `3 <= n <= 100_000`
- `0 <= nums[i] <= 100_000`
