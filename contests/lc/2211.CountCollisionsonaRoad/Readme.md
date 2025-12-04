# [2211. Count Collisions on a Road](https://leetcode.com/problems/count-collisions-on-a-road/description)

There are `n` cars on an infinitely long road, numbered `0` to `n - 1` from left to right. Each car starts at a unique point.

You are given a 0-indexed string `directions` of length `n` where `directions[i]` is:
- `L` if the `i`th car moves left
- `R` if the `i`th car moves right
- `S` if the `i`th car stays still

All moving cars share the same speed. Collisions update the total as follows:
- Two cars moving in opposite directions collide: +2 collisions
- Moving car hits a stationary car: +1 collision
- After colliding, involved cars become stationary and stay put

Return the total number of collisions that will occur.

## Examples

**Example 1**
```
Input: directions = "RLRSLL"
Output: 5
```
Collisions:
- Cars 0 and 1 collide (opposite directions): total = 2
- Cars 2 and 3 collide (car 3 stationary): total = 3
- Cars 3 and 4 collide (car 3 stationary): total = 4
- Cars 4 and 5 collide after 4 becomes stationary: total = 5

**Example 2**
```
Input: directions = "LLRR"
Output: 0
```
No cars collide.

## Constraints
- `1 <= directions.length <= 10^5`
- `directions[i]` is `L`, `R`, or `S`
