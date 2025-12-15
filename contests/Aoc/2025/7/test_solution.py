import subprocess
import sys
from pathlib import Path


def run_py(args):
    return subprocess.check_output([sys.executable, "solution.py", *args], text=True).strip()


def write_grid(tmp_path: Path, lines: list[str]) -> Path:
    p = tmp_path / "grid.txt"
    p.write_text("\n".join(lines))
    return p


def test_empty_file(tmp_path: Path):
    grid = write_grid(tmp_path, [])
    out1 = run_py([str(grid)])
    out2 = run_py(["--part2", str(grid)])
    assert out1 == "0"
    assert out2 == "0"


def test_no_start(tmp_path: Path):
    grid = write_grid(tmp_path, ["....", ".^.."])
    assert run_py([str(grid)]) == "0"
    assert run_py(["--part2", str(grid)]) == "0"


def test_single_split(tmp_path: Path):
    grid = write_grid(tmp_path, ["..S..", "..^.."])
    # Part1: 1 split; Part2: two timelines (left/right)
    assert run_py([str(grid)]) == "1"
    assert run_py(["--part2", str(grid)]) == "2"


def test_edge_falloff(tmp_path: Path):
    grid = write_grid(tmp_path, ["S", "^"])
    # Split sends beams off-grid; no surviving timeline
    assert run_py([str(grid)]) == "1"
    assert run_py(["--part2", str(grid)]) == "0"


def test_sample_input():
    base = Path(__file__).parent
    out1 = run_py([])
    out2 = run_py(["--part2"])
    assert out1 == "21"
    assert out2 == "40"
