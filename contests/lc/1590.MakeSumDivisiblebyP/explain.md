# 1590. Make Sum Divisible by P — Explanation

## Intuition
We want to remove the smallest contiguous subarray so the remaining sum is divisible by `p`. If the total sum is already divisible by `p`, answer is `0`. Otherwise, we need to find a subarray whose sum modulo `p` equals `target = totalSum % p`; removing it fixes the remainder.

This reduces to finding the shortest pair of prefix sums `prefix[j]` and `prefix[i]` (with `j > i`) where:
```
(prefix[j] - prefix[i]) % p == target
```
Rearranging:
```
prefix[i] % p == (prefix[j] - target + p) % p
```

## Algorithm
1. Compute `target = sum(nums) % p`; if `target == 0`, return `0`.
2. Maintain running prefix modulo `p` and a map `remainder -> earliest index` (initialize `{0: -1}` for empty prefix).
3. For each index `i`:
   - Update `prefix = (prefix + nums[i]) % p`.
   - The needed earlier remainder is `need = (prefix - target + p) % p`.
   - If `need` exists in the map, update best length with `i - index_of_need`.
   - Record/overwrite `map[prefix] = i` to keep the latest index for this remainder.
4. If the best length is smaller than `n`, return it; otherwise `-1` (cannot remove the whole array).

## Correctness
- The map stores the earliest position for each remainder seen so far. Any subarray sum from `(index_of_need+1) ... i` has remainder `target`, so removing it makes the remaining sum divisible by `p`.
- Scanning left to right ensures we consider all subarrays exactly once and track the minimum length.
- Rejecting `best == n` avoids removing the whole array, as required.

## Complexity
- Time: `O(n)` — single pass with constant-time map operations.
- Space: `O(min(n, p))` — at most one entry per possible remainder encountered.
