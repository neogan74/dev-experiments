# [2043. Simple Bank System](https://leetcode.com/problems/simple-bank-system/description/)

## Overview
You are building an automation layer for a bank that supports three transaction types: `transfer`, `deposit`, and `withdraw`. Accounts are numbered from `1` to `n`, and each has an initial balance described by the 0-indexed array `balance`, where account `i + 1` starts with `balance[i]`.

## Valid Transaction Criteria
- Account numbers referenced in any operation must be between `1` and `n`.
- Withdrawn or transferred funds must not exceed the balance of the source account.

## API Specification
- `Bank(long[] balance)`: Initialize the banking system with the provided balances.
- `boolean transfer(int account1, int account2, long money)`: Move `money` from `account1` to `account2`. Returns `true` on success; otherwise `false`.
- `boolean deposit(int account, long money)`: Add `money` to `account`. Returns `true` on success; otherwise `false`.
- `boolean withdraw(int account, long money)`: Remove `money` from `account`. Returns `true` on success; otherwise `false`.

## Example
```text
Input:
["Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"]
[[[10, 100, 20, 50, 30]], [3, 10], [5, 1, 20], [5, 20], [3, 4, 15], [10, 50]]

Output:
[null, true, true, true, false, false]
```

Explanation:
- `Bank bank = new Bank([10, 100, 20, 50, 30]);`
- `bank.withdraw(3, 10);   // true`
- `bank.transfer(5, 1, 20); // true`
- `bank.deposit(5, 20);     // true`
- `bank.transfer(3, 4, 15); // false`
- `bank.withdraw(10, 50);   // false`

## Constraints
- `n == balance.length`
- `1 <= n, account, account1, account2 <= 10^5`
- `0 <= balance[i], money <= 10^12`
- At most `10^4` calls will be made to each of `transfer`, `deposit`, and `withdraw`.
