# [3461. Check If Digits Are Equal in String After Operations I](https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/description)

## Problem

You are given a string `s` consisting of digits. Perform the following operation repeatedly until the string has exactly two digits:

- For each pair of consecutive digits in `s`, starting from the first digit, compute `(s[i] + s[i + 1]) % 10`.
- Replace `s` with the sequence of the newly calculated digits, preserving their order.

Return `true` if the final two digits in `s` are the same; otherwise, return `false`.

## Examples

### Example 1

**Input:** `s = "3902"`  
**Output:** `true`  
**Explanation:**

```
3902  -> 292  (2, 9, 2)
292   -> 11   (1, 1)
```

The final digits are both `1`, so the result is `true`.

### Example 2

**Input:** `s = "34789"`  
**Output:** `false`  
**Explanation:**

```
34789 -> 7157 -> 862 -> 48
```

The final digits are `4` and `8`, which differ, so the result is `false`.

## Constraints

- `3 <= s.length <= 100`
- `s` consists of only digits.
