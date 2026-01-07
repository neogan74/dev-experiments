# [1411. Number of Ways to Paint N x 3 Grid](https://leetcode.com/problems/number-of-ways-to-paint-n-3-grid/description)

You have a grid of size n x 3 and you want to paint each cell of the grid with exactly one of the three colors: Red, Yellow, or Green while making sure that no two adjacent cells have the same color (i.e., no two cells that share vertical or horizontal sides have the same color).

Given n the number of rows of the grid, return the number of ways you can paint this grid. As the answer may grow large, the answer must be computed modulo 10^9 + 7.

## Examples

Example 1:

![img.png](img.png)

Input: n = 1
Output: 12
Explanation: There are 12 possible way to paint the grid as shown.

Example 2:

Input: n = 5000
Output: 30228214

## Constraints

1 <= n <= 5000

## Approach

Each row has 3 cells with no horizontal adjacency conflicts. There are two pattern types:

- ABA: first and third colors are the same, middle is different (6 ways).
- ABC: all three colors are different (6 ways).

Let a be the number of ways where the current row is ABA, and b be the number of ways where the current row is ABC.
For each next row:

- From ABA: 3 ways to stay ABA, 2 ways to go ABC.
- From ABC: 2 ways to go ABA, 2 ways to stay ABC.

So:

- a' = 3a + 2b
- b' = 2a + 2b

Initialize a = 6, b = 6 for n = 1 and iterate n - 1 times.

## Complexity

Time: O(n)  
Space: O(1)
