# [2141. Maximum Running Time of N Computers](https://leetcode.com/problems/maximum-running-time-of-n-computers/description)

## Problem
You have `n` computers and a 0-indexed array `batteries`, where `batteries[i]` is the minutes of charge available in battery `i`. At most one battery can be inserted into each computer at a time, but batteries can be swapped instantly at integer minutes and cannot be recharged. Return the maximum number of minutes all `n` computers can run simultaneously.

## Examples

![two batteries swapped](image.png)
```
Input: n = 2, batteries = [3,3,3]
Output: 4
```

![four small batteries](image-1.png)
```
Input: n = 2, batteries = [1,1,1,1]
Output: 2
```

## Approach
- Observe that only the total usable energy matters; we can think of redistributing charge freely up to each battery's capacity.
- Binary search on the answer `t`: check if the computers can all run for `t` minutes by summing `min(battery, t)` across batteries; it is feasible when this sum is at least `n * t`.
- Search bounds: `low = 0`, `high = total_energy / n` (integer division) because we cannot exceed the average energy per computer.

## Complexity
- Time: `O(m log(total/n))` where `m = len(batteries)` (each check is linear).
- Space: `O(1)` beyond the input.
