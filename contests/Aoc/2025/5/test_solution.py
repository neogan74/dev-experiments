import pathlib
import sys
import unittest

sys.path.append(str(pathlib.Path(__file__).resolve().parent))

import solution


EXAMPLE_INPUT = """\
3-5
10-14
16-20
12-18

1
5
8
11
17
32
"""


def _parse_example() -> tuple[list[tuple[int, int]], list[int]]:
    ranges: list[tuple[int, int]] = []
    ids: list[int] = []
    reading_ranges = True
    for raw in EXAMPLE_INPUT.strip().splitlines():
        line = raw.strip()
        if not line:
            reading_ranges = False
            continue
        if reading_ranges:
            ranges.append(solution.parse_range(line))
        else:
            ids.append(int(line))
    return ranges, ids


class SolutionTests(unittest.TestCase):
    def test_example_part1(self) -> None:
        ranges, ids = _parse_example()
        self.assertEqual(solution.count_fresh(ids, ranges), 3)

    def test_example_part2(self) -> None:
        ranges, _ = _parse_example()
        self.assertEqual(solution.count_fresh_coverage(ranges), 14)

    def test_merge_overlapping_and_adjacent(self) -> None:
        ranges = [(4, 6), (1, 3), (10, 12), (6, 9)]
        # Combined ranges: 1-9 (9 IDs) and 10-12 (3 IDs) -> 12 total.
        self.assertEqual(solution.count_fresh_coverage(ranges), 12)


if __name__ == "__main__":
    unittest.main()
