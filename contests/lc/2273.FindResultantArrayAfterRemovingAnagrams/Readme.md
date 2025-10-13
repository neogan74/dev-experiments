# [2273. Find Resultant Array After Removing Anagrams](https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/description)

## Problem
You are given a 0-indexed string array `words`, where `words[i]` consists of lowercase English letters.

In one operation, select any index `i` such that `0 < i < words.length` and `words[i - 1]` and `words[i]` are anagrams, and delete `words[i]` from `words`. Keep performing this operation as long as you can select an index that satisfies the conditions.

Return `words` after performing all operations. It can be shown that selecting the indices for each operation in any arbitrary order will lead to the same result.

An anagram is a word or phrase formed by rearranging the letters of a different word or phrase using all the original letters exactly once. For example, `dacb` is an anagram of `abdc`.

## Examples
### Example 1
```text
Input:  words = ["abba","baba","bbaa","cd","cd"]
Output: ["abba","cd"]
```
Explanation:
- Since `words[2] = "bbaa"` and `words[1] = "baba"` are anagrams, delete `words[2]`.
  Now `words = ["abba","baba","cd","cd"]`.
- Since `words[1] = "baba"` and `words[0] = "abba"` are anagrams, delete `words[1]`.
  Now `words = ["abba","cd","cd"]`.
- Since `words[2] = "cd"` and `words[1] = "cd"` are anagrams, delete `words[2]`.
  Now `words = ["abba","cd"]`.
  No further operations are possible, so `["abba","cd"]` is the final answer.

### Example 2
```text
Input:  words = ["a","b","c","d","e"]
Output: ["a","b","c","d","e"]
```
Explanation: No two adjacent strings in `words` are anagrams, so no operations are performed.

## Constraints
- `1 <= words.length <= 100`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters
