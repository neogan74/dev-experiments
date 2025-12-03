import unittest

from solution import Solution


class ClosestPrimesTests(unittest.TestCase):
    def setUp(self) -> None:
        self.solver = Solution()

    def test_sample_gap_two(self):
        self.assertEqual(self.solver.closestPrimes(10, 19), [11, 13])

    def test_not_enough_primes(self):
        self.assertEqual(self.solver.closestPrimes(4, 6), [-1, -1])

    def test_single_prime_only(self):
        self.assertEqual(self.solver.closestPrimes(1, 2), [-1, -1])

    def test_tightest_gap_one(self):
        self.assertEqual(self.solver.closestPrimes(2, 3), [2, 3])

    def test_choose_earliest_on_tie(self):
        self.assertEqual(self.solver.closestPrimes(2, 11), [2, 3])


if __name__ == "__main__":
    unittest.main()
