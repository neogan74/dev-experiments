# [2523. Closest Prime Numbers in Range](https://leetcode.com/problems/closest-prime-numbers-in-range/description)

## Problem
Find two primes `num1 < num2` inside `[left, right]` with the smallest gap. If multiple pairs share the gap, pick the one with the smaller `num1`. Return `[-1, -1]` when fewer than two primes exist.

Constraints: `1 <= left <= right <= 1e6`.

## Approach
- Sieve all primes up to `right` once (Sieve of Eratosthenes).
- Scan numbers from `max(left, 2)` to `right`, tracking the previous prime seen.
- Update the best pair when the current gap is smaller than the best gap found so far.
- If fewer than two primes are found, return `[-1, -1]`.

## Complexity
- Time: `O(right log log right)` for the sieve and linear scan.
- Space: `O(right)` for the composite flag array.
