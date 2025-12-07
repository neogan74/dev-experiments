import argparse
from typing import List, Tuple


def parse_range(line: str) -> Tuple[int, int]:
    parts = line.strip().split("-")
    if len(parts) != 2:
        raise ValueError(f"invalid range {line!r}")
    start = int(parts[0])
    end = int(parts[1])
    if start > end:
        start, end = end, start
    return start, end


def load_input(path: str) -> Tuple[List[Tuple[int, int]], List[int]]:
    ranges: List[Tuple[int, int]] = []
    ids: List[int] = []
    reading_ranges = True

    with open(path, "r", encoding="utf-8") as f:
        for raw_line in f:
            line = raw_line.strip()
            if line == "":
                reading_ranges = False
                continue

            if reading_ranges:
                ranges.append(parse_range(line))
            else:
                ids.append(int(line))

    return ranges, ids


def count_fresh(ids: List[int], ranges: List[Tuple[int, int]]) -> int:
    total = 0
    for ingredient_id in ids:
        for start, end in ranges:
            if start <= ingredient_id <= end:
                total += 1
                break
    return total


def count_fresh_coverage(ranges: List[Tuple[int, int]]) -> int:
    if not ranges:
        return 0

    sorted_ranges = sorted(ranges, key=lambda r: r[0])
    merged_start, merged_end = sorted_ranges[0]
    coverage = 0

    for start, end in sorted_ranges[1:]:
        if start <= merged_end + 1:
            if end > merged_end:
                merged_end = end
            continue

        coverage += merged_end - merged_start + 1
        merged_start, merged_end = start, end

    coverage += merged_end - merged_start + 1
    return coverage


def main() -> None:
    parser = argparse.ArgumentParser(description="Day 5: Cafeteria (Advent of Code 2025)")
    parser.add_argument("path", nargs="?", default="input.txt", help="input file path")
    parser.add_argument("--part2", action="store_true", help="count total fresh IDs from ranges only")
    args = parser.parse_args()

    ranges, ids = load_input(args.path)
    if args.part2:
        result = count_fresh_coverage(ranges)
    else:
        result = count_fresh(ids, ranges)
    print(result)


if __name__ == "__main__":
    main()
