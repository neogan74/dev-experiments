# [1975. Maximum Matrix Sum](https://leetcode.com/problems/maximum-matrix-sum/description)

## Problem
You are given an `n x n` integer matrix. You can perform the following operation any
number of times:

- Choose any two adjacent elements and multiply each by `-1`.
- Two elements are adjacent if they share a border.

Return the maximum possible sum of all elements.

## Examples

### Example 1
![alt text](image-1.png)

Input: `matrix = [[1,-1],[-1,1]]`  
Output: `4`  
Explanation: We can follow the following steps to reach sum equals 4:
- Multiply the 2 elements in the first row by `-1`.
- Multiply the 2 elements in the first column by `-1`.

### Example 2
![alt text](image.png)

Input: `matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]`  
Output: `16`  
Explanation: We can follow the following step to reach sum equals 16:
- Multiply the 2 last elements in the second row by `-1`.

## Constraints

- `n == matrix.length == matrix[i].length`
- `2 <= n <= 250`
- `-10^5 <= matrix[i][j] <= 10^5`
