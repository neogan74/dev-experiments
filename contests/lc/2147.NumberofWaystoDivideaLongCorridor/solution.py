from typing import List


MOD = 1_000_000_007


def numberOfWays(corridor: str) -> int:
    """Return the number of ways to divide the corridor."""
    seats: List[int] = [i for i, ch in enumerate(corridor) if ch == "S"]

    if len(seats) == 0 or len(seats) % 2 == 1:
        return 0

    ways = 1
    for i in range(2, len(seats), 2):
        gap = seats[i] - seats[i - 1]
        ways = (ways * gap) % MOD

    return ways
