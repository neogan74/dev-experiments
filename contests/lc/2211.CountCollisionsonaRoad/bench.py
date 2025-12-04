"""
Lightweight benchmarks for solution.py and solution2.py.

Run:
  python bench.py
"""
from time import perf_counter

import solution
import solution2


CASES = [
    "RLRSLL",
    "LLRR",
    "S",
    "L",
    "R",
    "SRRS",
    "RRSLL",
    "RRSLLSRRSLRSLLRSRSLLRRSRLSLRRSLLSRS",
]


def run_bench(name, solver_cls, iterations=100_000):
    sol = solver_cls()
    # Warm up and correctness check.
    for directions in CASES:
        sol.countCollisions(directions)

    start = perf_counter()
    for _ in range(iterations):
        for directions in CASES:
            sol.countCollisions(directions)
    duration = perf_counter() - start
    total_calls = iterations * len(CASES)
    print(f"{name:10} | {total_calls:9} calls | {duration:7.3f}s | {duration / total_calls * 1e6:7.3f} us/call")


def main():
    run_bench("solution", solution.Solution)
    run_bench("solution2", solution2.Solution)


if __name__ == "__main__":
    main()
