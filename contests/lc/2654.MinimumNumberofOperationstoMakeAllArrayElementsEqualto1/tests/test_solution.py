from pathlib import Path
import sys

PROJECT_ROOT = Path(__file__).resolve().parents[1]
if str(PROJECT_ROOT) not in sys.path:
    sys.path.insert(0, str(PROJECT_ROOT))

from solution import Solution


def test_example_one():
    nums = [2, 6, 3, 4]
    assert Solution().minOperations(nums) == 4


def test_example_two():
    nums = [2, 10, 6, 14]
    assert Solution().minOperations(nums) == -1


def test_existing_ones_reduce_fast():
    nums = [1, 2, 3, 4]
    assert Solution().minOperations(nums) == 3


def test_all_ones_needs_no_operations():
    nums = [1, 1, 1]
    assert Solution().minOperations(nums) == 0


def test_needs_subarray_scan_for_gcd_one():
    nums = [6, 10, 15]
    assert Solution().minOperations(nums) == 4
