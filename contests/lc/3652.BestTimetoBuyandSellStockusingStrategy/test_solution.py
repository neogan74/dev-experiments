import unittest

from solution import Solution


class TestMaxProfit(unittest.TestCase):
    def test_variants(self):
        sol = Solution()
        cases = [
            # Example: modifying the first window boosts profit.
            ([4, 2, 8], [-1, 0, 1], 2, 10),
            # Example: best choice is leaving the strategy untouched.
            ([5, 4, 3], [1, 1, 0], 2, 9),
            # Best window is in the middle; flips the last element to a sell.
            ([1, 2, 3, 4], [-1, -1, -1, -1], 2, 1),
            # Whole-array modification where gains come from the new sells.
            ([5, 1, 3, 4], [-1, 1, 0, 1], 4, 7),
        ]

        for prices, strategy, k, expected in cases:
            with self.subTest(prices=prices, strategy=strategy, k=k):
                self.assertEqual(sol.maxProfit(prices, strategy, k), expected)
                self.assertEqual(sol.maxProfit2(prices, strategy, k), expected)


if __name__ == "__main__":
    unittest.main()
