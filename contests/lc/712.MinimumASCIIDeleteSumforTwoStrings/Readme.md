# [712. Minimum ASCII Delete Sum for Two Strings](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/description)

Given two strings `s1` and `s2`, return the lowest ASCII sum of deleted characters to make two strings equal.

## Examples

Example 1:

```
Input: s1 = "sea", s2 = "eat"
Output: 231
Explanation: Deleting "s" from "sea" adds 115. Deleting "t" from "eat" adds 116.
The total is 231.
```

Example 2:

```
Input: s1 = "delete", s2 = "leet"
Output: 403
Explanation: Delete "dee" from "delete" (100 + 101 + 101) and "e" from "leet" (101).
The total is 403.
```

## Constraints

- `1 <= s1.length, s2.length <= 1000`
- `s1` and `s2` consist of lowercase English letters.
