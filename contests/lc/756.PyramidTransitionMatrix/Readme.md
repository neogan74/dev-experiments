# [756. Pyramid Transition Matrix](https://leetcode.com/problems/pyramid-transition-matrix/description)

## Problem

You are stacking blocks to form a pyramid. Each block has a color, represented by a
single letter. Each row of blocks contains one less block than the row beneath it and
is centered on top.

Only specific triangular patterns are allowed. A triangular pattern consists of a
single block stacked on top of two blocks. The patterns are given as a list of
three-letter strings `allowed`, where the first two characters represent the left and
right bottom blocks and the third character is the top block.

For example, `ABC` represents a `C` block stacked on top of `A` (left) and `B` (right).
This is different from `BAC`, where `B` is on the left bottom and `A` is on the right
bottom.

You start with a bottom row of blocks `bottom`, given as a single string, that you
must use as the base of the pyramid.

Given `bottom` and `allowed`, return `true` if you can build the pyramid all the way
to the top such that every triangular pattern in the pyramid is in `allowed`, or
`false` otherwise.

## Examples

Example 1:

```
Input: bottom = "BCD", allowed = ["BCC","CDE","CEA","FFF"]
Output: true
Explanation: The allowed triangular patterns are shown on the right.
Starting from the bottom (level 3), we can build "CE" on level 2 and then build "A"
on level 1. There are three triangular patterns in the pyramid, which are "BCC",
"CDE", and "CEA". All are allowed.
```

Example 2:

```
Input: bottom = "AAAA", allowed = ["AAB","AAC","BCD","BBE","DEF"]
Output: false
Explanation: The allowed triangular patterns are shown on the right.
Starting from the bottom (level 4), there are multiple ways to build level 3, but
trying all the possibilities, you always get stuck before building level 1.
```

## Constraints

- `2 <= bottom.length <= 6`
- `0 <= allowed.length <= 216`
- `allowed[i].length == 3`
- The letters in all input strings are from `{'A', 'B', 'C', 'D', 'E', 'F'}`.
- All values of `allowed` are unique.
