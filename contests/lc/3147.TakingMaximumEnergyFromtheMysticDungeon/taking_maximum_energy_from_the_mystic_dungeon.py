from typing import List


def max_energy(energy: List[int], k: int) -> int:
    """Return the maximum energy obtainable by starting at any magician index."""
    n = len(energy)
    if n == 0:
        return 0

    totals = [0] * n
    best = float("-inf")

    for i in range(n - 1, -1, -1):
        jump = i + k
        totals[i] = energy[i] + (totals[jump] if jump < n else 0)
        if totals[i] > best:
            best = totals[i]

    return best


if __name__ == "__main__":  # pragma: no cover
    import sys
    import json

    if len(sys.argv) != 2:
        raise SystemExit("Usage: python taking_maximum_energy_from_the_mystic_dungeon.py '<energy,k>'")

    payload = json.loads(sys.argv[1])
    print(max_energy(payload["energy"], payload["k"]))
