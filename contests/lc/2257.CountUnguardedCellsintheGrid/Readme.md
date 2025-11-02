# 2257. Count Unguarded Cells in the Grid

You are given two integers `m` and `n` representing the dimensions of an `m x n` grid.

There are two sets of special cells:
- `guards[i] = [rowi, coli]` marks the location of a guard.
- `walls[j] = [rowj, colj]` marks the location of a wall.

Guards watch their row and column in the four orthogonal directions (up, down, left, right) until their view meets either the grid boundary, another guard, or a wall. A cell is considered **guarded** if it lies in the line of sight of at least one guard. Walls always block vision, and the guard or wall cells themselves are not counted as guarded.

Return the number of cells in the grid that are neither guarded nor occupied by a wall or guard—these cells are termed **unguarded**.

## Example 1

![alt text](image.png)

> **Input**: m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
>
> **Output**: 7
>
> **Explanation**: The guarded and unguarded cells are shown in red and green respectively in the above diagram.
There are a total of 7 unguarded cells, so we return 7.

# Example 2

![ex2](image-1.png)

> **Input**: m = 3, n = 3, guards = [[1,1]], walls = [[0,1],[1,0],[2,1],[1,2]]
>
> **Output**: 4
>
> **Explanation**: The unguarded cells are shown in green in the above diagram.
>
> There are a total of 4 unguarded cells, so we return 4.

## Constraints

- `1 <= m, n <= 105`
- `2 <= m * n <= 105`
- `1 <= guards.length, walls.length <= 5 * 104`
- `2 <= guards.length + walls.length <= m * n`
- `guards[i].length == walls[j].length == 2`
- `0 <= rowi, rowj < m`
- `0 <= coli, colj < n`
- All the positions in `guards` and `walls` are unique.
