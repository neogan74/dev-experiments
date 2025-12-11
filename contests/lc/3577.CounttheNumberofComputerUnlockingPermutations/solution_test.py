import unittest

from solution import Solution


class TestCountUnlockingPermutations(unittest.TestCase):
    def setUp(self) -> None:
        self.solver = Solution()

    def test_example_valid(self):
        self.assertEqual(self.solver.countUnlockingPermutations([1, 2, 3]), 2)

    def test_example_invalid(self):
        self.assertEqual(self.solver.countUnlockingPermutations([3, 3, 3, 4, 4, 4]), 0)

    def test_min_size_valid(self):
        self.assertEqual(self.solver.countUnlockingPermutations([1, 2]), 1)

    def test_root_not_unique_minimum(self):
        self.assertEqual(self.solver.countUnlockingPermutations([2, 1, 3]), 0)

    def test_factorial_result(self):
        self.assertEqual(self.solver.countUnlockingPermutations([1, 2, 5, 7, 8]), 24)  # (5-1)!

    def test_moderate_size(self):
        self.assertEqual(
            self.solver.countUnlockingPermutations([1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11]),
            3_628_800,  # 10!
        )


if __name__ == "__main__":
    unittest.main()
