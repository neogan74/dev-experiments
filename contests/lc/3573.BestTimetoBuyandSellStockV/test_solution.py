import unittest

from solution import Solution


class TestSolution(unittest.TestCase):
    def test_examples(self):
        sol = Solution()
        self.assertEqual(sol.maximumProfit([1, 7, 9, 8, 2], 2), 14)
        self.assertEqual(sol.maximumProfit([12, 16, 19, 19, 8, 1, 19, 13, 9], 3), 36)

    def test_no_transactions(self):
        sol = Solution()
        self.assertEqual(sol.maximumProfit([5, 4, 3], 0), 0)

    def test_single_price(self):
        sol = Solution()
        self.assertEqual(sol.maximumProfit([5], 1), 0)


if __name__ == "__main__":
    unittest.main()
