# 3494. Find the Minimum Amount of Time to Brew Potions

[LeetCode problem statement](https://leetcode.com/problems/find-the-minimum-amount-of-time-to-brew-potions/description)

## Problem

You are given two integer arrays `skill` and `mana` of lengths `n` and `m`, respectively.

In a laboratory, `n` wizards must brew `m` potions in order. Each potion has a mana capacity `mana[j]` and must pass through all the wizards sequentially to be brewed properly. The time taken by the `i`-th wizard on the `j`-th potion is `time[i][j] = skill[i] * mana[j]`.

Because the brewing process is delicate, a potion must be passed to the next wizard immediately after the current wizard completes their work. The start times need to be synchronized so that each wizard begins working on a potion exactly when it arrives.

Return the minimum amount of time required for all potions to be brewed properly.

## Examples

### Example 1

- Input: `skill = [1, 5, 2, 4]`, `mana = [5, 1, 4, 2]`
- Output: `110`
- Explanation:

| Potion | Start | Wizard 0 done | Wizard 1 done | Wizard 2 done | Wizard 3 done |
| ------ | ----- | ------------- | ------------- | ------------- | ------------- |
| 0      | 0     | 5             | 30            | 40            | 60            |
| 1      | 52    | 53            | 58            | 60            | 64            |
| 2      | 54    | 58            | 78            | 86            | 102           |
| 3      | 86    | 88            | 98            | 102           | 110           |

Wizard 0 cannot begin potion 1 at `t = 50`. Wizard 2 would finish potion 1 at `t = 58`, but wizard 3 is still busy with potion 0 until `t = 60`, so potion 1 cannot advance.

### Example 2

- Input: `skill = [1, 1, 1]`, `mana = [1, 1, 1]`
- Output: `5`
- Explanation:

```
Potion 0: start t = 0, finish t = 3
Potion 1: start t = 1, finish t = 4
Potion 2: start t = 2, finish t = 5
```

### Example 3

- Input: `skill = [1, 2, 3, 4]`, `mana = [1, 2]`
- Output: `21`

## Constraints

- `n == skill.length`
- `m == mana.length`
- `1 <= n, m <= 5000`
- `1 <= mana[j], skill[i] <= 5000`
