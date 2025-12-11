# [3577. Count the Number of Computer Unlocking Permutations](https://leetcode.com/problems/count-the-number-of-computer-unlocking-permutations/description)

## Problem
We need the number of permutations of `[0, 1, …, n-1]` that represent a valid unlocking order.

- Initially only computer `0` is unlocked.
- To unlock `i > 0`, there must already be an unlocked computer `j < i` with `complexity[j] < complexity[i]`.
- Return the count modulo `1e9+7`.

## Observation
Computer `1` can only rely on computer `0`, so we must have `complexity[0] < complexity[1]`. By extension, computer `0` must be strictly smaller than every other complexity; otherwise some index would have no strictly lower, earlier unlocked computer.

If `complexity[0]` is the unique minimum:
- Computer `0` is valid at position `0` of every permutation.
- For any `i > 0`, `0 < i` and `complexity[0] < complexity[i]`, so computer `i` can be unlocked immediately after `0`.
- All remaining computers can then appear in any order.

So the total number of valid permutations is simply `(n-1)!` when `complexity[0]` is strictly smaller than all other entries; otherwise it is `0`.

## Why this works
- Necessity: If some `complexity[k] <= complexity[0]`, then computer `k` cannot be unlocked after `0` (needs smaller complexity) nor before `0` (label constraint `j < i`), so no permutation exists.
- Sufficiency: When `complexity[0]` is strictly smaller than all others, every `i > 0` has a valid predecessor `0`. After unlocking `0`, the remaining `n-1` computers are unrestricted and can appear in any order, yielding exactly `(n-1)!` permutations.

## Algorithm
1. Check whether `complexity[0]` is strictly less than every `complexity[i]` (`i > 0`). If not, return `0`.
2. Compute `(n-1)! % MOD` with `MOD = 1_000_000_007`.

## Complexity
- Time: `O(n)` to verify the minimum and compute the factorial.
- Space: `O(1)`.
