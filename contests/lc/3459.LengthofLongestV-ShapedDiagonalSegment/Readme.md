# [3459. Length of Longest V-Shaped Diagonal Segment](https://leetcode.com/problems/length-of-longest-v-shaped-diagonal-segment/description)

You are given an `n x m` integer matrix `grid` where each cell is `0`, `1`, or `2`. Return the length of the longest V-shaped diagonal segment; if none exists, return `0`.

## V-shaped segment rules
- Starts at a cell with `1`.
- Follows the repeating sequence `2, 0, 2, 0, ...` after the starting `1`.
- Moves diagonally (↘️, ↙️, ↗️, ↖️) and may make **at most one clockwise 90° turn** to another diagonal while continuing the sequence.

## Examples
- **Example 1**  
  Input: `[[2,2,1,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]`  
  Output: `5`  
  Path: `(0,2) → (1,3) → (2,4) ↷ (3,3) → (4,2)`  
  ![diagram](image.png)

- **Example 2**  
  Input: `[[2,2,2,2,2],[2,0,2,2,0],[2,0,1,1,0],[1,0,2,2,2],[2,0,0,2,2]]`  
  Output: `4`  
  Path: `(2,3) → (3,2) ↷ (2,1) → (1,0)`  
  ![diagram](image-1.png)

- **Example 3**  
  Input: `[[1,2,2,2,2],[2,2,2,2,0],[2,0,0,0,0],[0,0,2,2,2],[2,0,0,2,0]]`  
  Output: `5`  
  Path: `(0,0) → (1,1) → (2,2) → (3,3) → (4,4)`  
  ![diagram](image-2.png)

- **Example 4**  
  Input: `[[1]]`  
  Output: `1`  
  Path: `(0,0)`

## Constraints
- `n == grid.length`
- `m == grid[i].length`
- `1 <= n, m <= 500`
- `grid[i][j] ∈ {0, 1, 2}`
