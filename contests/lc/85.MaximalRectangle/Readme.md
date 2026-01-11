# [85. Maximal Rectangle](https://leetcode.com/problems/maximal-rectangle/description)

Level: Hard

Given a rows x cols binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

## Example 1

![Example diagram](image.png)

Input: `matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]`  
Output: `6`  
Explanation: The maximal rectangle is shown in the above picture.

## Example 2

Input: `matrix = [["0"]]`  
Output: `0`

## Example 3

Input: `matrix = [["1"]]`  
Output: `1`

## Constraints

- `rows == matrix.length`
- `cols == matrix[i].length`
- `1 <= rows, cols <= 200`
- `matrix[i][j]` is `'0'` or `'1'`

## Approach

- Treat each row as the base of a histogram of consecutive 1's above it.
- For each row, update heights and compute the largest rectangle in the histogram with a monotonic stack.
- Time: `O(rows * cols)`, Space: `O(cols)`.
