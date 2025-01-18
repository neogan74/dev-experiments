# [1368. Minimum Cost to Make at Least One Valid Path in a Grid](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description)

## Problem
You are given an `m x n` grid. Each cell contains a direction pointing to the next cell you should visit:

- `1`: right `(i, j) -> (i, j + 1)`
- `2`: left `(i, j) -> (i, j - 1)`
- `3`: down `(i, j) -> (i + 1, j)`
- `4`: up `(i, j) -> (i - 1, j)`

Notes:
- Some arrows may point outside the grid.
- Start at the top-left cell `(0, 0)` and want to reach the bottom-right cell `(m - 1, n - 1)`.
- A *valid path* follows the arrows (not necessarily the shortest).
- You may change a cell's arrow once at cost `1`.

Return the minimum cost to ensure at least one valid path exists.

## Examples

### Example 1
![grid](image-2.png)

```
Input:  grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]
Output: 3
```
Path: `(0,0) -> (0,1) -> (0,2) -> (0,3)` change to down (`+1`) -> `(1,3) -> (1,2) -> (1,1) -> (1,0)` change to down (`+1`) -> `(2,0) -> (2,1) -> (2,2) -> (2,3)` change to down (`+1`) -> `(3,3)`. Total cost = 3.

### Example 2
![grid](image-1.png)

```
Input:  grid = [[1,1,3],[3,2,2],[1,1,4]]
Output: 0
```
You can already follow the arrows from `(0,0)` to `(2,2)`.

### Example 3
![grid](image.png)

```
Input:  grid = [[1,2],[4,3]]
Output: 1
```

## Constraints
- `m == grid.length`
- `n == grid[i].length`
- `1 <= m, n <= 100`
- `1 <= grid[i][j] <= 4`
