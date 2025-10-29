# [2043. Simple Bank System](https://leetcode.com/problems/simple-bank-system/description)

## Problem

You are building a program for a bank to automate incoming transactions (`transfer`, `deposit`, and `withdraw`). The bank has `n` accounts numbered from `1` to `n`. The initial balance of each account is stored in a 0-indexed integer array `balance`, where the `(i + 1)`-th account has an initial balance of `balance[i]`.

Execute every valid transaction. A transaction is valid if:

- The provided account numbers are between `1` and `n`.
- The amount withdrawn or transferred does not exceed the balance of the source account.

Implement the `Bank` class with the following operations:

- `Bank(long[] balance)`: Initialize the bank using the provided balances.
- `boolean transfer(int account1, int account2, long money)`: Transfer `money` dollars from `account1` to `account2`. Return `true` if the transaction succeeds, otherwise `false`.
- `boolean deposit(int account, long money)`: Deposit `money` dollars into `account`. Return `true` if the transaction succeeds, otherwise `false`.
- `boolean withdraw(int account, long money)`: Withdraw `money` dollars from `account`. Return `true` if the transaction succeeds, otherwise `false`.

## Example

**Input**

```text
["Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"]
[[[10, 100, 20, 50, 30]], [3, 10], [5, 1, 20], [5, 20], [3, 4, 15], [10, 50]]
```

**Output**

```text
[null, true, true, true, false, false]
```

**Explanation**

```text
Bank bank = new Bank([10, 100, 20, 50, 30]);
bank.withdraw(3, 10);    // true: account 3 has $20, so withdrawing $10 is valid.
bank.transfer(5, 1, 20); // true: account 5 has $30, so transferring $20 is valid.
bank.deposit(5, 20);     // true: depositing $20 into account 5 is valid.
bank.transfer(3, 4, 15); // false: account 3 has $10, so transferring $15 is invalid.
bank.withdraw(10, 50);   // false: account 10 does not exist.
```

## Constraints

- `n == balance.length`
- `1 <= n, account, account1, account2 <= 10^5`
- `0 <= balance[i], money <= 10^12`
- At most `10^4` calls will be made to each function.
