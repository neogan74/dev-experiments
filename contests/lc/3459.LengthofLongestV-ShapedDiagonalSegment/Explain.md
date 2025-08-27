# Algorithm Explanation (Go)

## Idea
We need the longest diagonal path that:
- starts on a cell with value `1`;
- continues with the alternating pattern `2, 0, 2, 0, ...`;
- moves diagonally and may make **one clockwise 90° turn** to another diagonal.

We can start a search from every `1` in every diagonal direction. Depth-first search with memoization lets us extend the path while respecting the pattern and the single-turn rule.

## Directions
`dirs := []int{1, 1, -1, -1, 1}` encodes four diagonal steps:
- `k=0` → `(1, 1)` (down-right)
- `k=1` → `(1, -1)` (down-left)
- `k=2` → `(-1, -1)` (up-left)
- `k=3` → `(-1, 1)` (up-right)

Next cell `(x, y)` is `(i + dirs[k], j + dirs[k+1])`.

## DFS state
`dfs(i, j, k, cnt)` returns the length of the longest valid continuation starting **after** `(i, j)`:
- `(i, j)`: current position.
- `k`: current diagonal direction index.
- `cnt`: `1` if we can still turn clockwise once; `0` otherwise.

### Expected value
If `(i, j)` is `1`, the next must be `2`; otherwise the pattern is `2 - grid[i][j]` to alternate between `2` and `0`.

### Transition
1. Move one step to `(x, y)` along direction `k`. If out of bounds or value mismatch, return `0`.
2. Continue straight: `res = dfs(x, y, k, cnt)`.
3. If `cnt > 0`, try a clockwise turn to direction `(k+1) % 4` once: `res = max(res, dfs(x, y, turnDir, 0))`.
4. Memoize and return `res + 1` (count current step).

Memo key: `[i][j][k][cnt]`. A zero value means “uncomputed”; positive values store best length from that state.

## Driver
Iterate all cells; when `grid[i][j] == 1`, try all four directions and take `1 + dfs(...)` to include the starting `1`. Track the maximum.

## Correctness sketch
- **Pattern maintenance:** Next expected value is derived from the previous cell (`2 - grid[i][j]`), so every step alternates `2` and `0` after the initial `1`.
- **Bounds and value check:** DFS stops if the move goes outside the grid or hits a wrong value, ensuring all counted paths are valid.
- **Single clockwise turn:** The flag `cnt` permits at most one transition to `(k+1) % 4`; after the turn `cnt` becomes `0`, preventing further turns.
- **Longest path:** From any state we take the maximum of straight and (single) turn continuations, so memoized results capture the best reachable length. Trying every `1` in every direction ensures the global maximum is found.

## Complexity
- States: `O(m * n * 4 * 2) = O(mn)` with memoization.
- Each state computes constant work, so **time** is `O(mn)`.
- **Space:** `O(mn)` for the memo table and recursion stack up to `O(m+n)` depth.
