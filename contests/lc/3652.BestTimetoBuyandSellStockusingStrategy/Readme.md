# [3652. Best Time to Buy and Sell Stock using Strategy](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-using-strategy/description)

## Problem
Given arrays `prices` and `strategy` of equal length and an even integer `k`:
- `strategy[i]` is `-1` (buy), `0` (hold), or `1` (sell).
- You may modify at most one window of length `k`: first `k/2` positions become `0`, last `k/2` positions become `1`.
- Profit is `sum(strategy[i] * prices[i])`.

Return the maximum achievable profit. Buying/selling capacity is unlimited.

## Examples
- `prices = [4,2,8], strategy = [-1,0,1], k = 2` → `10` (modify indices `[0,1]` to `[0,1]`).
- `prices = [5,4,3], strategy = [1,1,0], k = 2` → `9` (no modification needed).

## Approach
1. Compute the base profit `base = sum(strategy[i] * prices[i])`.
2. For every window of length `k`, calculate the profit delta if modified:
   - First half contributes `-strategy[i] * prices[i]` (becomes `0`).
   - Second half contributes `(1 - strategy[i]) * prices[i]` (becomes `1`).
3. Use prefix sums of `prices` and `strategy[i] * prices[i]` to get each window's delta in `O(1)`, track the maximum delta.
4. Answer is `base + max(0, bestDelta)`.

## Complexity
- Time: `O(n)` for a single pass over all windows.
- Space: `O(n)` for prefix sums (can be reduced to `O(1)` with a sliding window).

## Constraints
- `2 <= prices.length == strategy.length <= 1e5`
- `1 <= prices[i] <= 1e5`
- `-1 <= strategy[i] <= 1`
- `2 <= k <= prices.length`, `k` is even
