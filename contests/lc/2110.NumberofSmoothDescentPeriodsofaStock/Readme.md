# [2110. Number of Smooth Descent Periods of a Stock](https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/description)

## Problem
You are given an integer array `prices`, where `prices[i]` is the stock price on the `i`-th day.

A *smooth descent period* is a contiguous subarray where each day (except the first) is exactly `1` lower than the previous day. Single days always qualify. Return the total number of smooth descent periods.

## Examples
- `prices = [3,2,1,4]` → `7` (`[3] [2] [1] [4] [3,2] [2,1] [3,2,1]`)
- `prices = [8,6,7,7]` → `4` (`[8] [6] [7] [7]`)
- `prices = [1]` → `1`

## Approach
Scan once while tracking the length of the current smooth descent run. When the `prices[i] - prices[i-1] == -1` condition holds, extend the run; otherwise start a new run. Each run of length `k` contributes `k * (k + 1) / 2` descent periods, so add this amount when a run ends and after the final element.

## Complexity
- Time: `O(n)`
- Space: `O(1)`
