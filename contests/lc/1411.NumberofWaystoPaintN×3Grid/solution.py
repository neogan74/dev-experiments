def numOfWays(n: int) -> int:
    mod = 10**9 + 7
    a = b = 6
    for _ in range(2, n + 1):
        a, b = (a * 3 + b * 2) % mod, (a * 2 + b * 2) % mod
    return (a + b) % mod
