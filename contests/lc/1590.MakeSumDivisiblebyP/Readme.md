# [1590. Make Sum Divisible by P](https://leetcode.com/problems/make-sum-divisible-by-p/description)

Find the length of the smallest subarray to remove (cannot remove the whole array) so the remaining sum is divisible by `p`. Return `-1` if impossible.

## Approach
- Let `target = totalSum % p`. If `target == 0`, nothing needs removal.
- We need a subarray with sum `% p == target`. Using prefix sums modulo `p`, we look for two prefixes `i` and `j` where `(prefix[j] - prefix[i]) % p == target`.
- Track the earliest index of each prefix remainder in a hash map: remainder → index.
- At position `i` with `prefix`, the needed earlier remainder is `(prefix - target + p) % p`. When found, update the best length.
- Answer must be smaller than `n` (cannot delete the entire array).

## Complexity
- Time: `O(n)`
- Space: `O(p)` worst-case, bounded by `O(n)` for the map.
