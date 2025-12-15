import argparse
from pathlib import Path


def read_grid(path: str) -> list[str]:
    return Path(path).read_text().splitlines()


def find_start(grid: list[str]) -> tuple[int, int] | None:
    for r, row in enumerate(grid):
        c = row.find("S")
        if c != -1:
            return r, c
    return None


def count_splits(grid: list[str]) -> int:
    if not grid:
        return 0

    start = find_start(grid)
    if start is None:
        return 0

    srow, scol = start
    beams = {scol}
    splits = 0

    for r in range(srow + 1, len(grid)):
        next_beams: set[int] = set()
        row = grid[r]
        width = len(row)

        for c in beams:
            if c >= width:
                continue
            cell = row[c]
            if cell == "^":
                splits += 1
                if c > 0:
                    next_beams.add(c - 1)
                if c + 1 < width:
                    next_beams.add(c + 1)
                continue
            next_beams.add(c)

        beams = next_beams
        if not beams:
            break

    return splits


def count_timelines(grid: list[str]) -> int:
    if not grid:
        return 0

    start = find_start(grid)
    if start is None:
        return 0

    srow, scol = start
    beams: dict[int, int] = {scol: 1}

    for r in range(srow + 1, len(grid)):
        next_beams: dict[int, int] = {}
        row = grid[r]
        width = len(row)

        for c, cnt in beams.items():
            if c >= width:
                continue
            if row[c] == "^":
                for nc in (c - 1, c + 1):
                    if 0 <= nc < width:
                        next_beams[nc] = next_beams.get(nc, 0) + cnt
                continue
            next_beams[c] = next_beams.get(c, 0) + cnt

        beams = next_beams
        if not beams:
            break

    return sum(beams.values())


def main() -> None:
    parser = argparse.ArgumentParser(description="Tachyon manifold simulator")
    parser.add_argument("input", nargs="?", default="input.txt", help="path to input file")
    parser.add_argument(
        "--part2", action="store_true", help="compute timeline count (part 2)"
    )
    args = parser.parse_args()

    grid = read_grid(args.input)
    if args.part2:
        print(count_timelines(grid))
    else:
        print(count_splits(grid))


if __name__ == "__main__":
    main()
