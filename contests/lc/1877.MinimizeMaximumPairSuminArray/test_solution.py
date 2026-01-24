import unittest

from solution import Solution


class TestMinPairSum(unittest.TestCase):
    def setUp(self) -> None:
        self.solver = Solution()

    def test_example_one(self) -> None:
        self.assertEqual(self.solver.minPairSum([3, 5, 2, 3]), 7)

    def test_example_two(self) -> None:
        self.assertEqual(self.solver.minPairSum([3, 5, 4, 2, 4, 6]), 8)

    def test_sorted_pairs(self) -> None:
        self.assertEqual(self.solver.minPairSum([1, 2, 3, 4]), 5)

    def test_duplicate_values(self) -> None:
        self.assertEqual(self.solver.minPairSum([4, 4, 4, 4]), 8)

    def test_unbalanced_values(self) -> None:
        self.assertEqual(self.solver.minPairSum([1, 100000, 2, 99999]), 100001)


if __name__ == "__main__":
    unittest.main()
