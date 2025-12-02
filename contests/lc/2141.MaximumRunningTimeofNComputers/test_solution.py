import unittest

from solution import Solution


class MaxRunTimeTests(unittest.TestCase):
    def test_two_computers_equal_batteries(self):
        self.assertEqual(4, Solution().maxRunTime(2, [3, 3, 3]))

    def test_two_computers_small_batteries(self):
        self.assertEqual(2, Solution().maxRunTime(2, [1, 1, 1, 1]))

    def test_three_computers_mixed_batteries(self):
        self.assertEqual(8, Solution().maxRunTime(3, [10, 10, 3, 5]))

    def test_all_batteries_sufficient(self):
        self.assertEqual(6, Solution().maxRunTime(2, [5, 4, 3]))


if __name__ == "__main__":
    unittest.main()
