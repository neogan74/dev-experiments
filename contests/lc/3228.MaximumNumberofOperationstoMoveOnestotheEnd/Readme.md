# [3228. Maximum Number of Operations to Move Ones to the End](https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/description)

You are given a binary string `s`. The goal is to count the maximum number of operations that can be applied while following the rules below.

## Operation
- Choose an index `i` with `i + 1 < s.length` such that `s[i] == '1'` and `s[i + 1] == '0'`.
- Move `s[i]` to the right until it reaches the end of the string or is directly before another `'1'`.
  - Example: for `s = "010010"` and `i = 1`, the string becomes `"000110"`.

## Examples
**Example 1**
```
Input:  s = "1001101"
Output: 4
Explanation:
1. i = 0 → "0011101"
2. i = 4 → "0011011"
3. i = 3 → "0010111"
4. i = 2 → "0001111"
```

**Example 2**
```
Input:  s = "00111"
Output: 0
```

## Constraints
- `1 <= s.length <= 10^5`
- `s[i] ∈ {'0', '1'}` for every index `i`
