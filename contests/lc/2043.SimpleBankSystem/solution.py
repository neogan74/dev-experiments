from __future__ import annotations

from typing import List


class Bank:
    """Simulates bank accounts with transfer, deposit, and withdraw operations."""

    def __init__(self, balance: List[int]) -> None:
        # Copy to prevent callers from mutating the internal state.
        self._balance: List[int] = list(balance)

    def transfer(self, account1: int, account2: int, money: int) -> bool:
        if not self._valid_account(account1) or not self._valid_account(account2):
            return False
        idx1 = account1 - 1
        idx2 = account2 - 1
        if self._balance[idx1] < money:
            return False
        self._balance[idx1] -= money
        self._balance[idx2] += money
        return True

    def deposit(self, account: int, money: int) -> bool:
        if not self._valid_account(account):
            return False
        self._balance[account - 1] += money
        return True

    def withdraw(self, account: int, money: int) -> bool:
        if not self._valid_account(account):
            return False
        idx = account - 1
        if self._balance[idx] < money:
            return False
        self._balance[idx] -= money
        return True

    def _valid_account(self, account: int) -> bool:
        return 1 <= account <= len(self._balance)

    @property
    def balance(self) -> List[int]:
        return list(self._balance)
