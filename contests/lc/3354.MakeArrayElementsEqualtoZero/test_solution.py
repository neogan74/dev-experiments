import unittest

from solution import Solution


class TestCountValidSelections(unittest.TestCase):
    def setUp(self):
        self.solution = Solution()

    def test_example_one(self):
        self.assertEqual(self.solution.countValidSelections([1, 0, 2, 0, 3]), 2)

    def test_example_two(self):
        self.assertEqual(self.solution.countValidSelections([2, 3, 4, 0, 4, 1, 0]), 0)

    def test_single_zero(self):
        self.assertEqual(self.solution.countValidSelections([0]), 2)

    def test_zero_then_positive(self):
        self.assertEqual(self.solution.countValidSelections([0, 1]), 1)

    def test_positive_then_zero(self):
        self.assertEqual(self.solution.countValidSelections([1, 0]), 1)


if __name__ == "__main__":
    unittest.main()
