import unittest

from solution import Solution


class TestMinSubarray(unittest.TestCase):
    def setUp(self):
        self.solver = Solution()

    def test_examples(self):
        self.assertEqual(self.solver.minSubarray([3, 1, 4, 2], 6), 1)
        self.assertEqual(self.solver.minSubarray([6, 3, 5, 2], 9), 2)
        self.assertEqual(self.solver.minSubarray([1, 2, 3], 3), 0)

    def test_impossible(self):
        self.assertEqual(self.solver.minSubarray([1, 2, 3], 7), -1)

    def test_single_removal(self):
        self.assertEqual(self.solver.minSubarray([5, 7, 10], 3), 1)

    def test_large_values(self):
        self.assertEqual(self.solver.minSubarray([1_000_000_000, 2_000_000_000], 3), 0)


if __name__ == "__main__":
    unittest.main()
