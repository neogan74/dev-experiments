import argparse
import sys
from typing import List


def read_lines() -> List[str]:
    return [line.rstrip("\n") for line in sys.stdin]


def pad_lines(lines: List[str]) -> List[str]:
    width = max((len(line) for line in lines), default=0)
    return [line.ljust(width) for line in lines]


def solve(grid: List[str], part2: bool) -> int:
    if not grid:
        return 0

    height = len(grid)
    width = len(grid[0])
    op_row = height - 1

    separator = [all(grid[r][c] == " " for r in range(height)) for c in range(width)]
    has_separator = any(separator)

    def eval_block(start: int, end: int) -> int:
        op = next((ch for ch in grid[op_row][start:end] if ch != " "), None)
        if op not in {"+", "*"}:
            raise ValueError(f"invalid operator in columns {start}-{end}")

        numbers: List[int] = []
        if part2:
            for c in range(end - 1, start - 1, -1):
                digits = [grid[r][c] for r in range(op_row) if grid[r][c].isdigit()]
                if digits:
                    numbers.append(int("".join(digits)))
        else:
            for r in range(op_row):
                segment = grid[r][start:end].strip()
                if segment:
                    numbers.append(int(segment))

        if not numbers:
            raise ValueError(f"no numbers found in columns {start}-{end}")

        if op == "+":
            return sum(numbers)
        product = 1
        for n in numbers:
            product *= n
        return product

    if has_separator:
        total = 0
        c = 0
        while c < width:
            while c < width and separator[c]:
                c += 1
            start = c
            while c < width and not separator[c]:
                c += 1
            if start < c:
                total += eval_block(start, c)
        return total

    # Fallback: operator proximity.
    ops = [(idx, ch) for idx, ch in enumerate(grid[op_row]) if ch in {"+", "*"}]
    if not ops:
        raise ValueError("no operators found")

    def nearest_op(col: int) -> int:
        best_idx = 0
        best_dist = abs(col - ops[0][0])
        for i in range(1, len(ops)):
            dist = abs(col - ops[i][0])
            if dist < best_dist or (dist == best_dist and ops[i][0] > ops[best_idx][0]):
                best_idx = i
                best_dist = dist
        return best_idx

    buckets: List[List[int]] = [[] for _ in ops]

    if part2:
        for c in range(width):
            digits = [grid[r][c] for r in range(op_row) if grid[r][c].isdigit()]
            if not digits:
                continue
            value = int("".join(digits))
            buckets[nearest_op(c)].append(value)
    else:
        for r in range(op_row):
            row = grid[r]
            c = 0
            while c < width:
                if not row[c].isdigit():
                    c += 1
                    continue
                start = c
                while c < width and row[c].isdigit():
                    c += 1
                value = int(row[start:c])
                center = (start + c - 1) // 2
                buckets[nearest_op(center)].append(value)

    total = 0
    for (col, op), nums in zip(ops, buckets):
        if not nums:
            raise ValueError(f"no numbers found for operator at column {col}")
        if op == "+":
            total += sum(nums)
        else:
            prod = 1
            for n in nums:
                prod *= n
            total += prod
    return total


def main() -> None:
    parser = argparse.ArgumentParser(description="Solve Trash Compactor worksheet.")
    parser.add_argument("-part2", action="store_true", dest="part2", help="use part 2 rules (columns right-to-left)")
    args = parser.parse_args()

    grid = pad_lines(read_lines())
    try:
        result = solve(grid, args.part2)
    except Exception as exc:  # pylint: disable=broad-except
        print(f"error: {exc}", file=sys.stderr)
        sys.exit(1)
    print(result)


if __name__ == "__main__":
    main()
