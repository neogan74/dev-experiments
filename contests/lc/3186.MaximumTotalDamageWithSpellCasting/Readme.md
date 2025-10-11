# [3186. Maximum Total Damage With Spell Casting](https://leetcode.com/problems/maximum-total-damage-with-spell-casting/description)

## Problem
A magician has a collection of damaging spells. The damage dealt by each spell is given in the array `power`, and multiple spells may share the same damage value.

## Casting Rules
- Casting a spell with damage `power[i]` prevents casting any spell whose damage is `power[i] - 2`, `power[i] - 1`, `power[i] + 1`, or `power[i] + 2`.
- Each spell can be cast at most once per sequence.

## Goal
Determine the maximum total damage achievable while respecting the casting rules.

## Examples
**Example 1**
```text
Input:  power = [1, 1, 3, 4]
Output: 6
```
The maximum total damage of 6 is obtained by casting the spells with damage values 1, 1, and 4.

**Example 2**
```text
Input:  power = [7, 1, 6, 6]
Output: 13
```
The maximum total damage of 13 is obtained by casting the spells with damage values 1, 6, and 6.

## Constraints
- `1 <= power.length <= 10^5`
- `1 <= power[i] <= 10^9`
