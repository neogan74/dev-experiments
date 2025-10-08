# [2300. Successful Pairs of Spells and Potions](https://leetcode.com/problems/successful-pairs-of-spells-and-potions/description/)

## Problem

You are given two positive integer arrays `spells` and `potions`, of length `n` and `m` respectively, where `spells[i]` represents the strength of the `i`-th spell and `potions[j]` represents the strength of the `j`-th potion.

You are also given an integer `success`. A spell and potion pair is considered successful if the product of their strengths is at least `success`.

Return an integer array `pairs` of length `n` where `pairs[i]` is the number of potions that will form a successful pair with the `i`-th spell.

## Examples

### Example 1

```
Input:  spells = [5,1,3], potions = [1,2,3,4,5], success = 7
Output: [4,0,3]
```

Explanation:
- `0`-th spell: `5 * [1,2,3,4,5] = [5,10,15,20,25]`. Four pairs are successful.
- `1`-st spell: `1 * [1,2,3,4,5] = [1,2,3,4,5]`. Zero pairs are successful.
- `2`-nd spell: `3 * [1,2,3,4,5] = [3,6,9,12,15]`. Three pairs are successful.

### Example 2

```
Input:  spells = [3,1,2], potions = [8,5,8], success = 16
Output: [2,0,2]
```

Explanation:
- `0`-th spell: `3 * [8,5,8] = [24,15,24]`. Two pairs are successful.
- `1`-st spell: `1 * [8,5,8] = [8,5,8]`. Zero pairs are successful.
- `2`-nd spell: `2 * [8,5,8] = [16,10,16]`. Two pairs are successful.

## Constraints

- `n == spells.length`
- `m == potions.length`
- `1 <= n, m <= 10^5`
- `1 <= spells[i], potions[i] <= 10^5`
- `1 <= success <= 10^10`
