# [3531. Count Covered Buildings](https://leetcode.com/problems/count-covered-buildings/description)

## Problem
You are given an integer `n` for an `n x n` city grid and a list of unique building coordinates `buildings[i] = [x, y]`.

A building is **covered** if there is at least one other building in all four cardinal directions (same row left and right, same column above and below). Return the number of covered buildings.

## Examples
Example 1  
![example 1](image-2.png)  
Input: `n = 3, buildings = [[1,2],[2,2],[3,2],[2,1],[2,3]]`  
Output: `1`

Example 2  
![example 2](image-1.png)  
Input: `n = 3, buildings = [[1,1],[1,2],[2,1],[2,2]]`  
Output: `0`

Example 3  
![example 3](image.png)  
Input: `n = 5, buildings = [[1,3],[3,2],[3,3],[3,5],[5,3]]`  
Output: `1`

## Approach
For each row and column track the minimum and maximum coordinate that contains a building.  
For a building at `(r, c)`:
- It has a left neighbor if `c` is not the minimum column in its row.  
- It has a right neighbor if `c` is not the maximum column in its row.  
- It has an above neighbor if `r` is not the minimum row in its column.  
- It has a below neighbor if `r` is not the maximum row in its column.  
If all four conditions hold, the building is covered.

## Algorithm
1. Scan all buildings and update `(min_col, max_col)` per row and `(min_row, max_row)` per column.
2. Iterate through the buildings again and count those that are not equal to the min or max in both their row and column.

## Complexity
- Time: `O(m)` after the initial scans (`m = len(buildings)`), with only constant work per building aside from the two passes.  
- Space: `O(m)` for the row and column maps in the worst case.

## Solutions
- Python: `solution.py`
- Go: `solution.go`
