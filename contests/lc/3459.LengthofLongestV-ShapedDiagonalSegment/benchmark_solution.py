import time
from typing import List

from solution import Solution


def build_diagonal_grid(rows: int, cols: int) -> List[List[int]]:
    grid = [[0] * cols for _ in range(rows)]
    diag = min(rows, cols)
    for i in range(diag):
        if i == 0:
            grid[i][i] = 1
        elif i == 1:
            grid[i][i] = 2
        else:
            grid[i][i] = 0 if i % 2 == 0 else 2
    return grid


def benchmark(label: str, grid: List[List[int]], rounds: int = 5, repeat: int = 3) -> None:
    solver = Solution()
    durations = []
    for _ in range(repeat):
        start = time.perf_counter()
        for _ in range(rounds):
            solver.lenOfVDiagonal(grid)
        durations.append((time.perf_counter() - start) / rounds)
    best = min(durations)
    avg = sum(durations) / len(durations)
    print(f"{label}: best={best*1e3:.3f}ms  avg={avg*1e3:.3f}ms over {rounds} runs (repeat={repeat})")


if __name__ == "__main__":
    small = [
        [2, 2, 1, 2, 2],
        [2, 0, 2, 2, 0],
        [2, 0, 1, 1, 0],
        [1, 0, 2, 2, 2],
        [2, 0, 0, 2, 2],
    ]
    large = build_diagonal_grid(200, 200)

    benchmark("small", small, rounds=200)
    benchmark("large", large, rounds=10)
