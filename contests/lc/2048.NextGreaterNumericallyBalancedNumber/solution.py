from collections import Counter


def is_numerically_balanced(x: int) -> bool:
    """Return True when the decimal representation of x is numerically balanced."""
    counts = Counter(str(x))
    for digit, freq in counts.items():
        if freq != int(digit):
            return False
    return True


def next_beautiful_number(n: int) -> int:
    """Return the smallest numerically balanced number strictly greater than n."""
    candidate = n + 1
    while True:
        if is_numerically_balanced(candidate):
            return candidate
        candidate += 1
