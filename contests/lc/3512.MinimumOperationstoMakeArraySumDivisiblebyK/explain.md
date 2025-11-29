# Solution Explanation

## Intuition
Only the total sum matters. Every operation subtracts `1` from a single element, so it subtracts `1` from the whole sum. To reach the next smaller number that is divisible by `k`, we only need to remove `sum(nums) % k` from the total. Doing fewer is impossible (sum would still leave the same remainder), and doing more is unnecessary because the first divisible sum already appears after that many decrements.

## Algorithm
1. Compute `rem = sum(nums) % k`.
2. Return `rem`.

### Pseudocode
```
function minOperations(nums, k):
    total = sum(nums)
    return total mod k
```

## Correctness Proof
Let `S = sum(nums)` and `r = S mod k` (0 ≤ r < k).

- Each operation decreases `S` by exactly 1, so after `t` operations the sum is `S - t`.
- We need `S - t` to be divisible by `k`, i.e., `(S - t) mod k = 0`.
- This is equivalent to `t mod k = S mod k = r`.
- The smallest non-negative `t` satisfying `t mod k = r` is `t = r`.
Therefore the algorithm returns the minimum possible number of operations, and any smaller `t` cannot make the sum divisible by `k`. Hence the solution is correct. ▢

## Complexity Analysis
- Time: O(n) to compute the sum.
- Space: O(1) additional space.

## Example Trace
```
nums = [3, 9, 7], k = 5
S = 3 + 9 + 7 = 19
rem = 19 % 5 = 4
answer = 4
```

### Visual schema
```
sum = 19
19 % 5 = 4

19   18   17   16   15   ...
 |    |    |    |    |
 -1   -1   -1   -1   divisible by 5 here
      1    2    3    4 operations
```
After 4 decrements the total reaches 15, the first divisible-by-5 total, matching the computed answer.
