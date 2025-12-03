# [3625. Count Number of Trapezoids II](https://leetcode.com/problems/count-number-of-trapezoids-ii/description)

## Problem

You are given a 2D integer array `points` where `points[i] = [xi, yi]` represents the coordinates of the `i`-th point on the Cartesian plane. Return the number of unique trapezoids that can be formed by choosing any four distinct points from `points`.

A trapezoid is a convex quadrilateral with at least one pair of parallel sides. Two lines are parallel if and only if they have the same slope.

## Examples

**Example 1**

- Input: `points = [[-3,2],[3,0],[2,3],[3,2],[2,-3]]`
- Output: `2`

There are two distinct ways to pick four points that form a trapezoid:

1. `[-3,2], [2,3], [3,2], [2,-3]`
2. `[2,3], [3,2], [3,0], [2,-3]`

![Example 1 trapezoids](image.png)

**Example 2**

- Input: `points = [[0,0],[1,0],[0,1],[2,1]]`
- Output: `1`

![Example 2 trapezoid](image-1.png)

## Constraints

- `4 <= points.length <= 500`
- `-1000 <= xi, yi <= 1000`
- All points are pairwise distinct.
