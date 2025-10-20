# [1625. Lexicographically Smallest String After Applying Operations](https://leetcode.com/problems/lexicographically-smallest-string-after-applying-operations/description/)

You are given an even-length numeric string `s` and integers `a` and `b`. You may apply the following operations to `s` any number of times and in any order:

- Add `a` to every digit at an odd index (0-indexed), wrapping around after 9.    Example: `s = "3456"`, `a = 5` → `s = "3951"`.
- Rotate `s` to the right by `b` positions.    Example: `s = "3456"`, `b = 1` → `s = "6345"`.

Return the lexicographically smallest string obtainable after any sequence of the operations.

Two strings of equal length are compared lexicographically by locating the first differing position and comparing the digits at that index.

## Examples

### Example 1
```
Input:  s = "5525", a = 9, b = 2
Output: "2050"

Start:  "5525"
Rotate: "2555"
Add:    "2454"
Add:    "2353"
Rotate: "5323"
Add:    "5222"
Add:    "5121"
Rotate: "2151"
Add:    "2050"
```
No smaller string can be produced.

### Example 2
```
Input:  s = "74", a = 5, b = 1
Output: "24"

Start:  "74"
Rotate: "47"
Add:    "42"
Rotate: "24"
```
No smaller string can be produced.

### Example 3
```
Input:  s = "0011", a = 4, b = 2
Output: "0011"
```
All reachable strings are lexicographically larger.

## Constraints

- `2 <= s.length <= 100`
- `s.length` is even
- `s` consists only of digits `0-9`
- `1 <= a <= 9`
- `1 <= b <= s.length - 1`
