import argparse
import bisect
from typing import List, Tuple


def parse_ranges(line: str) -> Tuple[List[Tuple[int, int]], int]:
    ranges: List[Tuple[int, int]] = []
    max_end = 0
    for part in line.split(","):
        part = part.strip()
        if not part:
            continue
        if "-" not in part:
            raise ValueError(f"invalid range {part!r}")
        start_s, end_s = part.split("-", 1)
        start = int(start_s)
        end = int(end_s)
        if start > end:
            start, end = end, start
        ranges.append((start, end))
        if end > max_end:
            max_end = end
    return ranges, max_end


def generate_repeated_numbers(limit: int, max_repeat: int) -> List[int]:
    # Generate all numbers up to limit that are formed by repeating a digit
    # sequence k times, k >= 2. If max_repeat == 0, allow any k >= 2.
    seen = set()
    len_seq = 1
    while True:
        pow10 = 10 ** len_seq
        min_seq = pow10 // 10
        if min_seq == 0:
            len_seq += 1
            continue

        any_added = False
        repeat = 2
        while True:
            if max_repeat and repeat > max_repeat:
                break

            geom = (pow10 ** repeat - 1) // (pow10 - 1)  # 1 + pow + ... + pow^(k-1)
            min_val = min_seq * geom
            if min_val > limit:
                break

            max_seq = pow10 - 1
            end_seq = min(max_seq, limit // geom)
            if end_seq >= min_seq:
                any_added = True
                for seq in range(min_seq, end_seq + 1):
                    val = seq * geom
                    if val <= limit:
                        seen.add(val)
            repeat += 1

        # If even the smallest repeated number for this len_seq exceeds limit, stop.
        if not any_added and (min_seq * ((pow10 * pow10 - 1) // (pow10 - 1)) > limit):
            break
        len_seq += 1

    nums = sorted(seen)
    return nums


def build_prefix(nums: List[int]) -> List[int]:
    prefix = [0]
    for v in nums:
        prefix.append(prefix[-1] + v)
    return prefix


def solve(line: str, part2: bool) -> int:
    ranges, max_end = parse_ranges(line)
    max_repeat = 0 if part2 else 2
    nums = generate_repeated_numbers(max_end, max_repeat)
    prefix = build_prefix(nums)

    total = 0
    for start, end in ranges:
        l = bisect.bisect_left(nums, start)
        r = bisect.bisect_right(nums, end)
        total += prefix[r] - prefix[l]
    return total


def main() -> None:
    parser = argparse.ArgumentParser(description="Sum repeated-pattern IDs in ranges.")
    parser.add_argument("path", nargs="?", default="input.txt", help="input file path")
    parser.add_argument("--part2", action="store_true", help="use part two rules (repeat >= 2 times)")
    args = parser.parse_args()

    with open(args.path, "r", encoding="utf-8") as f:
        line = f.read().strip()
    if not line:
        print(0)
        return

    result = solve(line, args.part2)
    print(result)


if __name__ == "__main__":
    main()
