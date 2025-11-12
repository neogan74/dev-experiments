# [2654. Minimum Number of Operations to Make All Array Elements Equal to 1](https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/description)

## Problem
You are given a 0-indexed array `nums` consisting of positive integers. Repeatedly choose an index `i` such that `0 <= i < n - 1` and replace either `nums[i]` or `nums[i + 1]` with `gcd(nums[i], nums[i + 1])`. Return the minimum number of operations required to make every element equal to `1`, or `-1` if impossible.

Constraints: `2 <= n <= 50`, `1 <= nums[i] <= 10^6`.

## Examples
| Input | Output | Explanation |
| --- | --- | --- |
| `[2,6,3,4]` | `4` | Reduce the subarray `[3,4]` to `1`, then propagate the `1` outward in three more steps. |
| `[2,10,6,14]` | `-1` | Every `gcd` stays even, so `1` cannot be reached. |

## Approach
1. Count how many entries already equal `1`. If at least one exists, each remaining non-one can be made `1` in a single step by pairing with an adjacent `1`, so the answer is `n - count_ones`.
2. If the `gcd` of the full array is greater than `1`, no sequence of operations can create a `1`; return `-1`.
3. Otherwise, find the shortest subarray whose `gcd` is `1` by expanding all possible windows and tracking the minimum length `L`. Turning that segment into a single `1` costs `L - 1` operations, and then spreading the `1` to the rest of the array costs `n - 1` more, yielding `L + n - 2`.

Implementation lives in `solution.py`.

## Complexity
- Time: `O(n^2)` to scan all subarrays until a `gcd` of `1` is found.
- Space: `O(1)` auxiliary.
