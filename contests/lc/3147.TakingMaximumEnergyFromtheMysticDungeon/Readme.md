# [3147. Taking Maximum Energy From the Mystic Dungeon](https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/description)

## Problem
You are given an array `energy` and an integer `k`. Each index `i` represents a magician whose energy value `energy[i]` is absorbed when you visit them. After absorbing energy from magician `i`, you are instantly transported to magician `i + k`, and the process repeats until you teleport past the end of the array. You must select a starting magician and follow the teleportation rule, accumulating every visited magician's energy (positive or negative). Return the maximum total energy obtainable over all starting positions.

## Examples
- **Example 1**
  ```
  Input:  energy = [5, 2, -10, -5, 1], k = 3
  Output: 3
  ```
  Start at index 1 (value `2`), teleport to index 4 (value `1`), and gain `2 + 1 = 3`.

- **Example 2**
  ```
  Input:  energy = [-2, -3, -1], k = 2
  Output: -1
  ```
  Start at index 2 (value `-1`) to minimize the loss.

## Constraints
- `1 <= energy.length <= 10^5`
- `-1000 <= energy[i] <= 1000`
- `1 <= k <= energy.length - 1`
