# [2169. Count Operations to Obtain Zero](https://leetcode.com/problems/count-operations-to-obtain-zero/description)

## Problem
Given two non-negative integers `num1` and `num2`, repeatedly subtract the smaller (or equal) value from the larger value. Count how many operations are needed until either number becomes zero.

## Examples
**Example 1**
```
Input:  num1 = 2, num2 = 3
Output: 3
```
Operations:
1. `3 - 2 -> (2, 1)`
2. `2 - 1 -> (1, 1)`
3. `1 - 1 -> (0, 1)`

**Example 2**
```
Input:  num1 = 10, num2 = 10
Output: 1
```
Only one subtraction (`10 - 10`) makes one operand zero.

## Constraints
- `0 <= num1, num2 <= 10^5`

## Intuition
The process mirrors the subtraction-based version of the Euclidean algorithm for GCD. Each operation decreases the maximum of the two numbers, and the loop ends once one of them reaches zero.

## Algorithm
1. Initialize a counter `steps = 0`.
2. While both numbers are non-zero:
   - If `num1 >= num2`, subtract `num2` from `num1`; otherwise subtract `num1` from `num2`.
   - Increment `steps`.
3. Return `steps`.

### Reference Implementation
```python
class Solution:
    def countOperations(self, num1: int, num2: int) -> int:
        steps = 0
        while num1 and num2:
            if num1 >= num2:
                num1 -= num2
            else:
                num2 -= num1
            steps += 1
        return steps
```

## Complexity
- Time: `O(steps)` operations; in practice it is bounded by the subtraction-based GCD process.
- Space: `O(1)` additional memory.
