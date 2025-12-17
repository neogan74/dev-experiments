# [3573. Best Time to Buy and Sell Stock V](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-v/description)

You are given an array `prices` where `prices[i]` is the stock price on day `i`, and an integer `k`. You may perform at most `k` transactions. Each transaction is either:
- Normal: buy on day `i`, sell later on day `j > i`, earning `prices[j] - prices[i]`.
- Short: sell on day `i`, buy back later on day `j > i`, earning `prices[i] - prices[j]`.

Transactions cannot overlap; you must close the current position before opening another, and you cannot open a new one on the same day you close the previous.

Return the maximum total profit achievable with at most `k` transactions.

## Examples
- Input: `prices = [1,7,9,8,2]`, `k = 2`  
  Output: `14`  
  Explanation: Buy day 0, sell day 2 (`+8`); sell short day 3, buy back day 4 (`+6`).

- Input: `prices = [12,16,19,19,8,1,19,13,9]`, `k = 3`  
  Output: `36`

## Constraints
- `2 <= prices.length <= 10^3`
- `1 <= prices[i] <= 10^9`
- `1 <= k <= prices.length / 2`

## Approach
Use dynamic programming with three states per transaction count:
- `rest[t]`: max profit with `t` completed transactions and no open position.
- `long[t]`: max profit after opening a long position with `t` completed transactions.
- `short[t]`: max profit after opening a short position with `t` completed transactions.

Iterate prices once. From previous day’s states, try closing long/short positions (moving to `rest[t+1]`), or opening long/short positions from `rest[t]`. Each day uses only prior states, so we never open and close on the same day. The answer is `max(rest)`.

## Complexity
Time `O(n * k)`, space `O(k)`.
