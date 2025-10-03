from __future__ import annotations

from statistics import mean
from timeit import timeit

from water_bottle_2 import maxBottlesDrunk


def benchmark(iterations: int = 5, calls_per_iteration: int = 200_000) -> None:
    """Run simple timing benchmarks for representative inputs."""
    scenarios = (
        ("SmallExchange", 10, 3),
        ("MidExchange", 40, 6),
        ("LargeExchange", 100, 12),
    )

    print(f"Benchmarking maxBottlesDrunk with {iterations} runs, {calls_per_iteration} calls each")
    for name, num_bottles, num_exchange in scenarios:
        stmt = "maxBottlesDrunk(num_bottles, num_exchange)"
        setup = (
            "from water_bottle_2 import maxBottlesDrunk\n"
            f"num_bottles = {num_bottles}\nnum_exchange = {num_exchange}"
        )
        timings = [
            timeit(stmt=stmt, setup=setup, number=calls_per_iteration)
            for _ in range(iterations)
        ]
        avg_time = mean(timings) / calls_per_iteration
        print(f"{name:>14}: avg {avg_time:.3e}s per call | best {min(timings) / calls_per_iteration:.3e}s")


if __name__ == "__main__":
    benchmark()
