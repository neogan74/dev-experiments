import random
import unittest

from solution import Solution


def brute_force(nums, k):
    """Return maximum subarray sum where length is divisible by k using O(n^2)."""
    best = None
    for i in range(len(nums)):
        total = 0
        for j in range(i, len(nums)):
            total += nums[j]
            length = j - i + 1
            if length % k == 0:
                if best is None or total > best:
                    best = total
    return best


class TestSolution(unittest.TestCase):
    def setUp(self):
        self.solve = Solution().maxSubarraySum

    def test_examples(self):
        self.assertEqual(self.solve([1, 2], 1), 3)
        self.assertEqual(self.solve([-1, -2, -3, -4, -5], 4), -10)
        self.assertEqual(self.solve([-5, 1, 2, -3, 4], 2), 4)

    def test_all_negative(self):
        nums = [-5, -4, -3, -2, -1]
        self.assertEqual(self.solve(nums, 2), brute_force(nums, 2))

    def test_single_element(self):
        self.assertEqual(self.solve([7], 1), 7)

    def test_random_small_arrays(self):
        random.seed(0)
        for _ in range(100):
            n = random.randint(1, 8)
            k = random.randint(1, n)
            nums = [random.randint(-5, 5) for _ in range(n)]
            expected = brute_force(nums, k)
            self.assertEqual(self.solve(nums, k), expected)


if __name__ == "__main__":
    unittest.main()
