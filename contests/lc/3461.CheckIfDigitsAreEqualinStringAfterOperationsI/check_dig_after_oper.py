from __future__ import annotations


def has_same_digits(s: str) -> bool:
    """
    Return True if repeated pairwise modular sums reduce `s` to two equal digits.

    The algorithm mirrors the iterative process:
    - replace each adjacent pair by their modulo-10 sum
    - continue until two digits remain
    """
    digits = [ord(ch) - ord("0") for ch in s]
    for k in range(len(digits) - 1, 1, -1):
        for i in range(k):
            digits[i] = (digits[i] + digits[i + 1]) % 10
    return digits[0] == digits[1]
