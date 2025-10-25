# [2048. Next Greater Numerically Balanced Number](https://leetcode.com/problems/next-greater-numerically-balanced-number/)

## Problem
An integer `x` is numerically balanced if, for every digit `d` in the number `x`, the digit `d` occurs exactly `d` times in `x`.

Given an integer `n`, return the smallest numerically balanced number strictly greater than `n`.

## Examples
**Example 1**
```text
Input: n = 1
Output: 22
```
- The digit `2` occurs two times, satisfying the numerically balanced property.
- `22` is the smallest numerically balanced number strictly greater than `1`.

**Example 2**
```text
Input: n = 1000
Output: 1333
```
- The digit `1` occurs once and the digit `3` occurs three times.
- `1022` is invalid because the digit `0` would occur more than zero times.

**Example 3**
```text
Input: n = 3000
Output: 3133
```
- The digit `1` occurs once and the digit `3` occurs three times.
- `3133` is the smallest numerically balanced number strictly greater than `3000`.

## Constraints
- `0 <= n <= 10^6`
