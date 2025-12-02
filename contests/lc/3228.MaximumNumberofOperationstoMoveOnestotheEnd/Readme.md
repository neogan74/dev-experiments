# [3228. Maximum Number of Operations to Move Ones to the End](https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/description)

You are given a binary string `s`. An operation chooses an index `i` with `s[i] == '1'` and `s[i + 1] == '0'`, then slides that `'1'` to the right until it sits just before the next `'1'` or reaches the end. The task is to count the maximum number of operations that can be performed.

## Approach
- Group the string into runs of `1`s separated by runs of `0`s. Every time a run of `0`s starts immediately after a `'1'`, all `1`s seen so far will eventually cross that boundary exactly once, contributing that many operations.
- Scan left to right, keeping `ones` as the number of `1`s seen. When `s[i] == '0'` and `s[i - 1] == '1'`, add `ones` to the answer. Other zeros are already inside a zero-run and contribute nothing new.
- This directly counts the total number of times any `'1'` is moved in an optimal sequence.

## Complexity
- Time: `O(n)` — single pass.
- Space: `O(1)` — constant extras.
